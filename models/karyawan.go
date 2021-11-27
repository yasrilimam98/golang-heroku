package models

import (
	"net/http"

	"github.com/yasrilimam98/grb-restapi/database"
)

type Users struct {
	// Id                int       `json:"id"`
	// Nama              string    `json:"name`
	// Email             string    `json:"email"`
	// Email_verified_at time.Time `json:"email_verified_at"`
	// Password          string    `json:"password`
	// Remember_token    string    `json:"remember_token"`
	// Created_at        time.Time `json:"created_at"`
	// Updated_at        time.Time `json:"updated_at"`
	Id      int    `json:"id"`
	Nama    string `json:"nama`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchAllUsers() (Response, error) {
	var obj Users
	var arrobj []Users
	var res Response

	con := database.CreateCondb1()

	sqlStatement := "SELECT * FROM dt_karyawan"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj
	return res, nil
}

func StoreUsers(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	// v := validator.New()

	// peg := Karyawan{
	// 	Nama:    nama,
	// 	Alamat:  alamat,
	// 	Telepon: telepon,
	// }

	// err := v.Struct(peg)
	// if err != nil {
	// 	return res, err
	// }

	con := database.CreateCondb1()

	// sqlStatement := "INSERT users (name, email, email_verified_at, password, remember_token, created_at, update_at ) VALUES (?, ?, ?)"
	sqlStatement := "INSERT dt_karyawan (nama, alamat, telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateUsers(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := database.CreateCondb1()

	// sqlStatement := "UPDATE users SET name = ?, email = ?, email_verified_at = ?, password = ?, remember_token = ?, created_at = ?, updated_at = ? WHERE id = ?"
	sqlStatement := "UPDATE dt_karyawan SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUsers(id int) (Response, error) {
	var res Response

	con := database.CreateCondb1()

	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
