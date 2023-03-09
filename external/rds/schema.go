package rds

var createUserTable = `
	CREATE TABLE IF NOT EXISTS Users (
	    userId int NOT NULL AUTO_INCREMENT,
	    username varchar(20) NOT NULL,
		email varchar(100) NOT NULL UNIQUE,
	    PRIMARY KEY (userId)
	);
`

var createSessionsTable = `
	CREATE TABLE IF NOT EXISTS Sessions (
	    sessionId int NOT NULL AUTO_INCREMENT,
	    userId int NOT NULL,
	    createdAt Timestamp NOT NULL,
	    PRIMARY KEY (sessionId),
	    FOREIGN KEY (userId) REFERENCES Users(userId)
	);
`

var createOtpTable = `
	CREATE TABLE IF NOT EXISTS Otps (
	    otp int NOT NULL,
	    userId int NOT NULL,
	    createdAt Timestamp NOT NULL,
	    PRIMARY KEY (userId)
	);
`
