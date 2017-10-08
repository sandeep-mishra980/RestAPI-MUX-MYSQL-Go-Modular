package main

import (
	"com.app/usercontroller"
	"com.app/util"
	"fmt"
)

func main() { //start of main func
	fmt.Println("starting up")
	util.ConnectionStringMain()
	usercontroller.UserController()
	// dao.connectionString2()
	// conn:=util.connectionString1()
	//sandeepdao.dbConnectionFunc()
	//usercontroller.connectionString1()
} //end of main func
