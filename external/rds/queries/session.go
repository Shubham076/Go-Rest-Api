package queries

import (
	"BootCampT1/external/rds"
	"BootCampT1/models"
	"database/sql"
)

func InsertSession(session models.Session) (sql.Result, error) {
	q := "INSERT INTO Sessions (userId, sessionId, createdAt) VALUES (?, ?, ?)"
	return rds.Db.Exec(q, session.UserId, session.SessionId, session.CreatedAt)
}

func DeleteSession(sessionId int) (sql.Result, error) {
	return rds.Db.Exec("DELETE FROM Sessions where sessionId = ?", sessionId)
}
