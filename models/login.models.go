package models

import (
	"database/sql"
	"fmt"

	"github.com/yasrilimam98/grb-restapi/helpers"

	"github.com/yasrilimam98/grb-restapi/database"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj User
	var pwd string

	con := database.CreateCondb1()

	err := con.QueryRow("SELECT id,email,password FROM cms_users WHERE email = ?", email).Scan(
		&obj.Id, &obj.Email, &pwd,
	)

	// Cek baris null atau ada

	if err == sql.ErrNoRows {
		fmt.Println("Email Anda salah!")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	// jika tidak ada error

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Password Anda salah!")
		return false, err
	}

	// kalau bener
	return true, nil
}
