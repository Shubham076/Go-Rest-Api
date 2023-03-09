package queries

import (
	"BootCampT1/external/rds"
	"BootCampT1/logger"
	"BootCampT1/models"
	"database/sql"
	"math/big"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	row := rds.Db.QueryRow("SELECT userId, email, username FROM Users WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Email, &user.Username)
	if err != nil {
		logger.Error.Println(err)
		return nil, err
	}
	return &user, nil
}

func GetUserById(id big.Int) (*models.User, error) {
	var user models.User
	row := rds.Db.QueryRow("SELECT userId, email, username FROM Users WHERE userId = ?", id)
	err := row.Scan(&user.Id, &user.Email, &user.Username)
	if err != nil {
		logger.Error.Println("err")
		return nil, err
	}
	return &user, nil
}

func InsertUser(user models.User) (sql.Result, error) {
	q := "INSERT INTO Users (email, username) VALUES (?, ?)"
	return rds.Db.Exec(q, user.Email, user.Username)
}
