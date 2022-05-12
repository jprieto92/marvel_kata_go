package main

import (
	"github.com/jprieto92/marvel_kata_go/pkg/server"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Marvel Project microservice")
	log.Println("Listening ...")
	serv := &server.Server{DbInfo: server.NewMarvelDatabaseInfo()}
	log.Fatal(http.ListenAndServe(":8080", serv))
}
