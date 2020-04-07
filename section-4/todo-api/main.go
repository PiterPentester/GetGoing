package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GetGoing/section-4/todo-api/controller"
	"github.com/GetGoing/section-4/todo-api/model"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "localhost"
	port   = "3000"
	server = host + ":" + port
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("[*] Serving on " + server)
	log.Fatal(http.ListenAndServe(server, mux))
}
