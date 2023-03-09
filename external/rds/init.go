package rds

import (
	"BootCampT1/config"
	"BootCampT1/logger"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

// data source name(dsn) for mysql username:password@protocol(address)/dbname?param=value
func getDsn() string {
	conf := config.GetConfig()
	rds := conf.Rds
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true", rds.User, rds.Password, rds.Protocol, rds.Host, rds.DB)
	return dsn
}

func CreateTables() {
	tx := Db.MustBegin()
	tx.MustExec(createUserTable)
	tx.MustExec(createSessionsTable)
	tx.MustExec(createOtpTable)
	err := tx.Commit()
	if err != nil {
		logger.Error.Printf("Unable to create table %s", err.Error())
	}

	logger.Info.Println("Tables created successfully")
}

func Connect() {
	dsn := getDsn()
	var err error
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.Error.Printf("Can't connect to Aws Rds %w", err)
		panic("Can't connect to DB")
	}
	logger.Info.Printf("Connection to Aws Rds is successful")
}
