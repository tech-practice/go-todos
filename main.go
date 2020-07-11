package main

import (
	"fmt"
	"go-todos/config"
	"go-todos/database"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	db := database.ConnectDB(conf.Mongo)
	fmt.Println(db)
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)
}
