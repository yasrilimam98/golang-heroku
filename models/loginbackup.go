package models

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/yasrilimam98/grb-restapi/helpers"

// 	"github.com/yasrilimam98/grb-restapi/database"
// )

// type User struct {
// 	Id       int    `json:"id"`
// 	Username string `json:"username"`
// }

// func CheckLogin(username, password string) (bool, error) {
// 	var obj User
// 	var pwd string

// 	con := database.CreateCondb2()

// 	sqlStatement := "SELECT * FROM users WHERE username = ?"

// 	err := con.QueryRow(sqlStatement, username).Scan(
// 		&obj.Id, &obj.Username, &pwd,
// 	)

// 	// Cek baris null atau ada

// 	if err == sql.ErrNoRows {
// 		fmt.Println("Username not found")
// 		return false, err
// 	}

// 	if err != nil {
// 		fmt.Println("Query error")
// 		return false, err
// 	}

// 	// jika tidak ada error

// 	match, err := helpers.CheckPasswordHash(password, pwd)
// 	if !match {
// 		fmt.Println("Hash and password doesn't match.")
// 		return false, err
// 	}

// 	// kalau bener
// 	return true, nil
// }
