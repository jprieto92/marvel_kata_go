package comics

import (
	"encoding/json"
	"fmt"
	"github.com/jprieto92/marvel_kata_go/pkg/model"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	ErrBadUri            = "uri not valid"
	ErrWithServer        = "error when get comics from URI"
	ErrStatusCodeNoOk    = "statusCode received no ok"
	ErrBodyResponse      = "error when try to read response body"
	ErrGettingAllComics  = "error when try to get all comics"
	ErrUnmarshallingJson = "error when try unmarshalling comics"
)

//Processor saves comics dbendpoint URI
type Processor struct {
	url string
}

//NewComicProcessor validate comics dbendpoint URI and if its compliance return new processor
func NewComicProcessor(uri string) (*Processor, error) {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", ErrBadUri, err)
	}
	return &Processor{url: uri}, nil
}

//GetAllComics retrieves all comics information of a week
func (p *Processor) GetAllComics() ([]byte, error) {
	resp, err := http.Get(p.url)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("%v: %w", ErrWithServer, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v: %w", ErrStatusCodeNoOk, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", ErrBodyResponse, err)
	}

	return body, nil
}

//GetComicsPublishedInWeekUntilTime retrieves all comics published in a wek until a certain time
func (p *Processor) GetComicsPublishedInWeekUntilTime(t time.Time) (string, error) {
	response, err := p.GetAllComics()
	if err != nil {
		return "", fmt.Errorf("%v: %w", ErrGettingAllComics, err)
	}

	marvelDb := model.MarvelDb{}
	err = json.Unmarshal(response, &marvelDb)
	if err != nil {
		return "", fmt.Errorf("%v: %w", ErrUnmarshallingJson, err)
	}

	publishedComics, err := marvelDb.SearchComicsUntilTime(t)
	return publishedComics.EncodeComicListToJson()
}
