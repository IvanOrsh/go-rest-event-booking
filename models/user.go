package models

import "github.com/IvanOrsh/go-rest-event-booking/db"

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

	res, err := stmt.Exec(u.email, u.password)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	u.ID = userId
	return err
}
