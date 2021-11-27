package models

// import (
// 	"github.com/yasrilimam98/grb-restapi/database"
// )

// func FetchAllClient(regno int, serial string) ([]Client, error) {
// 	con := database.CreateCondb2()
// 	rows, err := con.Query("SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END , DATEDIFF(duedate,curdate()),dbname,dbport,localnetwork,localport,publicnetwork,publicport,name FROM mclient Where regno = ? AND serial = ?", regno, serial)

// 	if err != nil {
// 		return nil, err
// 	} else {
// 		clients := []Client{}
// 		for rows.Next() {
// 			// var regno int
// 			// var serial string
// 			// var duedate string
// 			// var name string
// 			// var dbname string
// 			// var dbport string
// 			// var localnetwork string
// 			// var localport string
// 			// var publicnetwork string
// 			// var publicport string
// 			var status1 string
// 			var sisa1 string
// 			var dbname1 string
// 			var dbport1 string
// 			var locnet1 string
// 			var locport1 string
// 			var pubnet1 string
// 			var pubport1 string
// 			var name1 string
// 			err2 := rows.Scan(&status1, &sisa1, &dbname1, &dbport1, &locnet1, &locport1, &pubnet1, &pubport1, &name1)
// 			if err2 != nil {
// 				return nil, err2
// 			} else {
// 				client := Client{regno, serial, status1, sisa1, dbname1, dbport1, locnet1, locport1, pubnet1, pubport1}
// 				clients = append(clients, client)
// 			}
// 		}
// 		return clients, nil
// 	}
// }


// Back up awal =================================BATAS========================================

// import (
// 	"net/http"

// 	"github.com/yasrilimam98/grb-restapi/database"
// )

// type Client struct {
// 	Id            int    `json:"id"`
// 	Regno         int    `json:"regno"`
// 	Serial        string `json:"serial"`
// 	Date          string `json:"duedate"`
// 	Name          string `json:"name"`
// 	DBname        string `json:"dbname"`
// 	DBport        string `json:"dbport"`
// 	LocalNetwork  string `json:"localnetwork"`
// 	LocalPort     string `json:"localport"`
// 	PublicNetwork string `json:"publicnetwork"`
// 	PublicPort    string `json:"publicport"`
// }

// func FetchAllClient(regno int, serial string) (Response, error) {
// 	var obj Client
// 	var arrobj []Client
// 	var res Response
// 	con := database.CreateCondb2()

// 	// sqlStatement := "SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END as statuslisensi, DATEDIFF(duedate,curdate()) as jmlhari,dbname,dbport,localnetwork,localport,publicnetwork,publicport,`name` FROM mclient Where regno = ? AND serial = ?",regno,serial)
// 	// rows, err := con.Query("SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END as id, DATEDIFF(duedate,curdate()) as serial,dbname,dbport,localnetwork,localport,publicnetwork,publicport,`name` FROM mclient Where regno = ? AND serial = ?", regno, serial)
// 	rows, err := con.Query("call getclientduedate(?, ?)", regno, serial)
// 	// rows, err := con.Query("SELECT id, regno, serial, duedate, name, dbname, dbport,localnetwork, localport, publicnetwork, publicport FROM mclient Where regno = ? AND serial = ?", regno, serial)
// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&obj.Id, &obj.Regno, &obj.Serial, &obj.Date, &obj.Name, &obj.DBname, &obj.DBport, &obj.LocalNetwork, &obj.LocalPort, &obj.PublicNetwork, &obj.PublicPort)
// 		if err != nil {
// 			return res, err
// 		}
// 		arrobj = append(arrobj, obj)
// 	}
// 	res.Status = http.StatusOK
// 	res.Message = "Succes"
// 	res.Data = arrobj
// 	return res,nil
// }



// package models

// import (
// 	"net/http"

// 	"github.com/yasrilimam98/grb-restapi/database"
// )

// // type Client struct {
// // 	Regno         int
// // 	Serial        string
// // 	Status        string
// // 	Name          string
// // 	DBname        string
// // 	DBport        string
// // 	LocalNetwork  string
// // 	LocalPort     string
// // 	PublicNetwork string
// // 	PublicPort    string
// // }
// type Client struct {
// 	Regno         int
// 	Serial        string
// 	Status        string
// 	SisaHari      string
// 	Name          string
// 	DBname        string
// 	DBport        string 
// 	LocalNetwork  string
// 	LocalPort     string
// 	PublicNetwork string
// 	PublicPort    string 
// }

// func FetchAllClient(regno int, serial string) (Response, error) {
// 	var obj Client
// 	var arrobj []Client
// 	var res Response
// 	con := database.CreateCondb2()

// 	rows, err := con.Query("SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END as id, DATEDIFF(duedate,curdate()) as serial,dbname,dbport,localnetwork,localport,publicnetwork,publicport,`name` FROM mclient Where regno = ? AND serial = ?", regno, serial)

// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	} else {
// 		for rows.Next() {
// 			err2 := rows.Scan(&obj.Status, &obj.SisaHari, &obj.Name, &obj.DBname, &obj.DBport, &obj.LocalNetwork, &obj.LocalPort, &obj.PublicNetwork, &obj.PublicPort)
// 			if err2 != nil {
// 				return res, err2
// 			} else {
// 				arrobj = append(arrobj, obj)
// 			}
// 		}
// 		res.Status = http.StatusOK
// 		res.Message = "Succes"
// 		res.Data = arrobj
// 		return res, nil
// 	}
// }

