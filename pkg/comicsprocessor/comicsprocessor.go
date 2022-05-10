package comicsprocessor

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type DictionaryErr string

var (
	ErrBadUri         = DictionaryErr("uri not valid")
	ErrWithServer     = DictionaryErr("error when get comics from URI")
	ErrStatusCodeNoOk = DictionaryErr("statusCode received no ok")
	ErrBodyResponse   = DictionaryErr("error when try to read response body")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

type Comic struct {
	Name  string
	Date  time.Time
	Price float64
	Url   string
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
		log.Println("Error ocurred when try to get commics", "URI:", p.url, "Error:", err)
		return "", ErrWithServer
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("No OK status code received ", "statusCode:", resp.StatusCode)
		return "", ErrStatusCodeNoOk
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error ocurred when try to read body response", "Error:", err)
		return "", ErrBodyResponse
	}

	return string(body), nil
}
