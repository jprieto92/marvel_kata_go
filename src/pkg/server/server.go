package server

import (
	"github.com/jprieto92/marvel_kata_go/src/pkg/comics"
	"log"
	"net/http"
	"time"
)

type server struct {
	uridb string
}

func NewServer(uridb string) *server {
	return &server{uridb: uridb}
}

func (s *server) Listen() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Println("Handling new request")
			processor, err := comics.NewComicProcessor(s.uridb)
			if err != nil {
				log.Println("Error occurred when try to create dbprocessor", "Err:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			result, err := processor.GetComicsPublishedInWeekUntilTime(time.Now())
			if err != nil {
				log.Println("Error occurred when try to get comics published in this week", "Err:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			log.Println("Comics published in this week retrieved successfully")
			if result == "{}" {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result))
			return
		}
		log.Println("Method not allowed", "Method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
