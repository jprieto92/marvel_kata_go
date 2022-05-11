package server

import (
	"fmt"
	"github.com/jprieto92/marvel_kata_go/pkg/comics"
	"github.com/jprieto92/marvel_kata_go/pkg/model"
	"log"
	"net/http"
	"time"
)

type DataBaseInfo interface {
	GetDatabaseUri() string
}

type MarvelDbInfo struct {
	dbUri string
}

func NewMarvelDatabaseInfo() *MarvelDbInfo {
	return &MarvelDbInfo{dbUri: model.MarvelDbUri}
}

func (db *MarvelDbInfo) GetDatabaseUri() string {
	return db.dbUri
}

type Server struct {
	DbInfo DataBaseInfo
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("Handling new request")
		processor, err := comics.NewComicProcessor(s.DbInfo.GetDatabaseUri())
		if err != nil {
			log.Println("Trying. Error when creating dbprocessor", "Err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		result, err := processor.GetComicsPublishedInWeekUntilTime(time.Now())
		if err != nil {
			log.Println("Trying. Error when getting published in this week", "Err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
			return
		}
		log.Println("Success. Comics published in this week retrieved")
		if result == "[]" {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintf(w, http.StatusText(http.StatusNoContent))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
		return
	}
	log.Println("Method not allowed", "Method:", r.Method)
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, http.StatusText(http.StatusMethodNotAllowed))
	return
}
