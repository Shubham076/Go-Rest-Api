package controllers

import (
	"BootCampT1/config"
	"BootCampT1/external/RabbitMq"
	"BootCampT1/external/RabbitMq/handlers"
	"BootCampT1/external/rds/queries"
	"BootCampT1/external/ses"
	"BootCampT1/logger"
	"BootCampT1/models"
	"BootCampT1/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	var user *models.User
	var err error

	conf := config.GetConfig()
	err = utils.UnmarshallAndValidate(&req, c)
	if err != nil {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	user, err = queries.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	if user == nil {
		msg := fmt.Sprintf("No user found with email: %s", req.Email)
		logger.Error.Printf(msg)
		utils.BadResponse(msg, c)
		return
	}

	//send otp
	otp := utils.GenerateOtp()
	err = queries.RemoveOldAndInsertNewOtpForUserId(models.Otp{Otp: otp, UserId: user.Id, CreatedAt: time.Now().UTC()})
	if err != nil {
		logger.Error.Printf("Unable to store otp for userId: %s", user.Id)
		utils.SendResponse("Failed to save user", c, http.StatusInternalServerError)
		return
	}
	utils.SendMail(handlers.SendMailRequest{
		EmailConf: ses.EmailConfig{
			Sender:    conf.Email.Sender,
			Recipient: []string{user.Email},
			Subject:   "One time password for secure login",
			HtmlBody:  `<h2>Otp: </h2>` + strconv.Itoa(otp),
			TextBody:  conf.Email.TextBody,
		},
	})

	utils.OKResponse("Successfully delivered otp to email address for secure login", c)
}

func SignUpHandler(c *gin.Context) {
	var req SignUpRequest
	var user *models.User
	var err error

	conf := config.GetConfig()
	err = utils.UnmarshallAndValidate(&req, c)
	if err != nil {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	user, err = queries.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	if user != nil {
		msg := fmt.Sprintf("User with %s already present in the DB", req.Email)
		logger.Error.Printf(msg)
		utils.BadResponse(msg, c)
		return
	}

	//create new user
	newUser := models.User{Email: req.Email, Username: req.Username}
	res, err := queries.InsertUser(newUser)

	if err != nil {
		logger.Error.Printf("Can't save user in the db, err: %w", err.Error())
		utils.SendResponse("Failed to save user", c, http.StatusInternalServerError)
		return
	}

	//generate otp
	otp := utils.GenerateOtp()
	userId, err := res.LastInsertId()
	if err != nil {
		logger.Error.Println("Unable to get userId, err: %s", err.Error())
		return
	}

	err = queries.RemoveOldAndInsertNewOtpForUserId(models.Otp{Otp: otp, UserId: userId, CreatedAt: time.Now().UTC()})
	if err != nil {
		logger.Error.Printf("Unable to store otp for userId: %s", userId)
		utils.SendResponse("Failed to save user", c, http.StatusInternalServerError)
		return
	}

	utils.SendMail(handlers.SendMailRequest{
		EmailConf: ses.EmailConfig{
			Sender:    conf.Email.Sender,
			Recipient: []string{newUser.Email},
			Subject:   "One time password to verify your email address",
			HtmlBody:  `<h2>Otp: </h2>` + strconv.Itoa(otp),
			TextBody:  conf.Email.TextBody,
		},
	})
	utils.OKResponse("Successfully delivered otp to email address, please verify the mail", c)
}

func Logouthandler(c *gin.Context) {
	var req LogoutRequest
	var err error
	err = utils.UnmarshallAndValidate(&req, c)
	if err != nil {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	err = RabbitMq.Push("DeleteSession", req)
	if err != nil {
		logger.Error.Printf("Logout failed for user: %v, err: %v", req.UserId, err)
		return
	}
}

func VerifyOtp(c *gin.Context) {
	var req VerifyOtpRequest
	var err error
	var otp *models.Otp
	conf := config.GetConfig()
	err = utils.UnmarshallAndValidate(&req, c)
	if err != nil {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	otp, err = queries.GetOtpByUserId(req.UserId)
	if err != nil || err == sql.ErrNoRows {
		logger.Error.Printf("Not able to find otp for userId: %v", req.UserId)
		utils.SendResponse("Unable to verify opt for the user", c, http.StatusInternalServerError)
		return
	}

	if otp.Otp != req.Otp {
		logger.Error.Printf("Invalid otp")
		utils.SendResponse("Invalid otp", c, http.StatusUnauthorized)
		return
	}

	cur := time.Now().Unix()
	prev := otp.CreatedAt.Unix()
	limit := conf.Otp.Expiry
	if cur-prev > limit {
		logger.Error.Println("Otp expired")
		utils.SendResponse("Otp expired", c, http.StatusUnauthorized)
		return
	}

	//delete verified otp
	_, err = queries.DeleteOtpsForUserId(req.UserId)
	if err != nil {
		logger.Error.Printf("Failed to delete verified otp for userId: %v, err: %v", req.UserId, err)
		utils.SendResponse("Unable to verify otp", c, http.StatusInternalServerError)
		return
	}

	//create a session
	session := models.Session{UserId: req.UserId, CreatedAt: time.Now().UTC()}
	_, err = queries.InsertSession(session)
	if err != nil {
		logger.Error.Printf("Unable to create session for userId: %v, err: %v", req.UserId, err)
		utils.SendResponse("Unable to verify otp", c, http.StatusInternalServerError)
		return
	}

	utils.OKResponse("success", c)
}

func GetUserData(c *gin.Context) {
	var req GetUserDetailsRequest
	var user *models.User
	var err error

	err = utils.UnmarshallAndValidate(&req, c)
	if err != nil {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	user, err = queries.GetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		logger.Error.Println(err)
		utils.BadResponse(err.Error(), c)
		return
	}

	if user == nil {
		msg := fmt.Sprintf("User with email: %s not found", req.Email)
		logger.Error.Printf(msg)
		utils.BadResponse("User not found", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
