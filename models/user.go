package models

import (
	"github.com/IvanOrsh/go-rest-event-booking/db"
	"github.com/IvanOrsh/go-rest-event-booking/utils"
)

type User struct {
	ID       int64
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (
			email, password
		)
		VALUES (
			?, ?
		)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	u.ID = userId
	return err
}
