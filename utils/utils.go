package utils

import (
	"BootCampT1/config"
	"BootCampT1/external/RabbitMq"
	"BootCampT1/external/RabbitMq/handlers"
	"BootCampT1/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func SendResponse(message string, c *gin.Context, status int) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

func BadResponse(message string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

func OKResponse(message string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func UnmarshallAndValidate(obj any, c *gin.Context) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return err
	}
	validate := validator.New()
	err = validate.Struct(obj)
	if err != nil {
		return err
	}
	return nil
}

func GenerateOtp() int {
	conf := config.GetConfig()
	max := conf.Otp.Max
	min := conf.Otp.Min
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(max-min) + min
	return otp
}

func SendMail(data handlers.SendMailRequest) {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		logger.Error.Printf(err.Error())
		return
	}

	err = RabbitMq.Push("Email", data)
	if err != nil {
		logger.Error.Printf("Unable to send mail due to err: %s", err.Error())
	}
}
