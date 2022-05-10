package comicsprocessor

import (
	"encoding/json"
	"github.com/jprieto92/marvel_kata_go/pkg/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type DictionaryErr string

var (
	ErrBadUri            = DictionaryErr("uri not valid")
	ErrWithServer        = DictionaryErr("error when get comics from URI")
	ErrStatusCodeNoOk    = DictionaryErr("statusCode received no ok")
	ErrBodyResponse      = DictionaryErr("error when try to read response body")
	ErrGettingAllComics  = DictionaryErr("error when try to get all comics")
	ErrUnmarshallingJson = DictionaryErr("error when try unmarshalling comics")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

type Processor struct {
	url string
}

func NewComicProcessor(uri string) (*Processor, error) {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		log.Println("Uri not valid", "Uri:", uri)
		return nil, ErrBadUri
	}
	return &Processor{url: uri}, nil
}

func (p *Processor) GetAllComics() (string, error) {
	resp, err := http.Get(p.url)
	defer resp.Body.Close()
	if err != nil {
		log.Println("Error occurred when try to get comics", "URI:", p.url, "Error:", err)
		return "", ErrWithServer
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("No OK status code received ", "statusCode:", resp.StatusCode)
		return "", ErrStatusCodeNoOk
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error occurred when try to read body response", "Error:", err)
		return "", ErrBodyResponse
	}

	return string(body), nil
}

func (p *Processor) GetComicsPublishedInWeekUntilTime(time time.Time) (string, error) {
	response, err := p.GetAllComics()
	if err != nil {
		log.Println("Error occurred when try to get all comics", "Error:", err)
		return "", ErrGettingAllComics
	}

	var marvelDb = model.MarvelDb{}
	err = json.Unmarshal([]byte(response), &marvelDb)
	if err != nil {
		log.Println("Error occurred when unmarshalling comics json", "Error:", err)
		return "", ErrUnmarshallingJson
	}

	for _, comic := range marvelDb.Data.Results {
		for _, date := range comic.Dates {
			if date.Type == "onsaleDate" {
				convertTimestampToDate(date.Date)
				// TO-DO : very date and create new comic element if
				// satisfies our requirement
			}
		}
	}

	return "", nil
}

func convertTimestampToDate(timestamp string) time.Time {

	return time.Now()
}
