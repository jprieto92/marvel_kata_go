package main

import (
	"fmt"
	"github.com/jprieto92/marvel_kata_go/pkg/comics"
	"github.com/jprieto92/marvel_kata_go/pkg/model"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Starting Marvel Project tool")
	processor, err := comics.NewComicProcessor(model.MarvelDbUri)
	if err != nil {
		log.Fatalf("Error occurred when try to create dbprocessor", "Err:", err)
		os.Exit(11)
	}

	result, err := processor.GetComicsPublishedInWeekUntilTime(time.Now())
	if err != nil {
		log.Fatalf("Error occurred when try to get comics published in this week", "Err:", err)
		os.Exit(12)
	}
	log.Println("Comics published in this week retrieved successfully:")
	fmt.Println(result)
	os.Exit(0)
}
