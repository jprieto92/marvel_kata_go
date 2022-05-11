package main

import (
	"github.com/jprieto92/marvel_kata_go/src/pkg/model"
	"github.com/jprieto92/marvel_kata_go/src/pkg/server"
	"log"
)

func main() {
	log.Println("Starting Marvel Project microservice")

	s := server.NewServer(model.MarvelDbUri)
	s.Listen()
}
