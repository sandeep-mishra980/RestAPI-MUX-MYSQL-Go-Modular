package usercontroller

import (
	"com.app/dao"
	_ "database/sql"
	_ "encoding/json"
	_ "flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "log"
	"net/http"
)

func UserController() {
	fmt.Println("In UserController")
	r := mux.NewRouter()
	r.HandleFunc("/users", dao.GetAllUsersFunc).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8888", nil)
}
