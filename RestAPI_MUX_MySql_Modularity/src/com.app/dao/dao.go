package dao

import (
	"com.app/databasestruct"
	"com.app/model"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
	"net/http"
)

var err error

type samUser model.User //typedef of struct

func GetAllUsersFunc(res http.ResponseWriter, req *http.Request) { //start of GetAllUserFunc
	//db,err =connectionString()//getting connection of database as a return
	db := databasestruct.DBCon
	fmt.Println("In GetAllUsersFunc")
	fmt.Println("Database", db)
	fmt.Println(err)
	// checkErr(err)
	fmt.Println("before--reaching in GetAllUsersFunc funcn")
	rows, err := db.Query("SELECT * FROM user ")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(rows)
	defer rows.Close()
	setDataRow := samUser{}
	results := []samUser{}
	for rows.Next() { //start of loop for fetching rows from db
		var id int
		var name, address string
		rows.Scan(&id, &name, &address)
		setDataRow.ID = id
		setDataRow.Name = name
		setDataRow.Address = address
		results = append(results, setDataRow)
	} //end of loop for fetching rows from db

	fmt.Println(results)
	//converting struct data into json
	b, err := json.Marshal(results) //will return data in byte formate
	if err != nil {
		fmt.Println(err)
		return
	}
	// string(b)-->will convert byte data into string but not will be formated
	fmt.Println(string(b)) //printing on console as formated bcz formate package using
	//json.NewEncoder(w).Encode(string(b))//not formated jason data as a  response from server
	fmt.Fprintf(res, string(b)) //formated json data as a response

	fmt.Println("after--reaching in homehandler funcn")
} //end of GetAllUserFunc
