package main

import (
	"fmt"
	"github.com/jprieto92/marvel_kata_go/pkg/comics"
	"github.com/jprieto92/marvel_kata_go/pkg/model"
	"log"
	"time"
)

func main() {
	log.Println("Starting Marvel Project tool")
	processor, err := comics.NewComicProcessor(model.MarvelDbUri)
	if err != nil {
		log.Fatalf("Trying. Error when creating dbprocessor", "Err:", err)
	}

	result, err := processor.GetComicsPublishedInWeekUntilTime(time.Now())
	if err != nil {
		log.Fatalf("Trying. Error when getting comics published in this week", "Err:", err)
	}
	log.Println("Comics published in this week retrieved successfully:")
	fmt.Println(result)
}
