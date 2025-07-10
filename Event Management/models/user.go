package models

import (
	"Event_Management/db"
	"Event_Management/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	// hashing the password
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	user, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := user.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = userId

	return err

}

func (u User) ValidateCredentials() error {
	query := `
SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		// invalid email
		return errors.New("Invalid email or password")
	}

	isValidPassword := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !isValidPassword {
		// invalid password
		return errors.New("Invalid email or password")
	}

	return nil
}
