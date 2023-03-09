package queries

import (
	"BootCampT1/external/rds"
	"BootCampT1/logger"
	"BootCampT1/models"
	"database/sql"
	"fmt"
)

func GetOtpByUserId(userId int64) (*models.Otp, error) {
	var otp models.Otp
	row := rds.Db.QueryRow("SELECT otp, createdAt FROM Otps WHERE userId = ?", userId)
	err := row.Scan(&otp.Otp, &otp.CreatedAt)
	if err != nil {
		logger.Error.Println(err)
		return nil, err
	}
	return &otp, nil
}

func InsertOtp(otp models.Otp) (sql.Result, error) {
	fmt.Println(otp.CreatedAt)
	q := "INSERT INTO Otps (otp, userId, createdAt) VALUES (?, ?, ?)"
	return rds.Db.Exec(q, otp.Otp, otp.UserId, otp.CreatedAt)
}

func DeleteOtpsForUserId(userId int64) (sql.Result, error) {
	return rds.Db.Exec("DELETE FROM Otps where userId = ?", userId)
}

func RemoveOldAndInsertNewOtpForUserId(otp models.Otp) error {
	_, err := DeleteOtpsForUserId(otp.UserId)
	if err != nil {
		logger.Error.Printf("Failed to delete records for userId: %v, err: %v", otp.UserId, err)
		return err
	}

	_, err = InsertOtp(otp)
	if err != nil {
		logger.Error.Printf("Failed the insert otp for userId: %v, err: %v", otp.UserId, err)
		return err
	}

	return nil
}
