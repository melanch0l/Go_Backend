package models

import (
	"errors"
	"example/restapi/db"
	"example/restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := ` INSERT INTO users(
	email,password)
	VAlUES(?,?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err //nil if work
}
func (u *User) Validate() error {
	query := `SELECT id,password
	FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var hashPassword string
	//retrive id,password from db
	err := row.Scan(&u.ID, &hashPassword)
	if err != nil {
		return errors.New("password invalid")
	}
	isValidPassword := utils.CheckPasswordHash(hashPassword, u.Password)
	if !isValidPassword {
		return errors.New("password invalid")
	}
	return nil

}
