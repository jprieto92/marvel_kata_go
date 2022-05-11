package model

import (
	"github.com/jprieto92/marvel_kata_go/src/pkg/utils"
	"time"
)

const MarvelDbUri = "https://gateway.marvel.com/v1/public/comics?dateDescriptor=thisWeek&limit=100&ts=987&apikey=97f295907072a970c5df30d73d1f3816&hash=abfa1c1d42a73a5eab042242335d805d"

// MarvelDb data type created using https://mholt.github.io/json-to-go/
type MarvelDb struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Results []struct {
			Title string `json:"title"`
			Urls  []Url  `json:"urls"`
			Dates []struct {
				Type string `json:"type"`
				Date string `json:"date"`
			} `json:"dates"`
			Prices []Price `json:"prices"`
		} `json:"results"`
	} `json:"data"`
}

type Url struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Price struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

//SearchComicsUntilTime search all comics published until a time in a marvelDB struct
func (db *MarvelDb) SearchComicsUntilTime(t time.Time) (ListComics, error) {
	publishedComics := ListComics{}

	for _, c := range db.Data.Results {
		for _, date := range c.Dates {
			if date.Type == "onsaleDate" {
				timestamp, err := utils.ConvertTimestampToDate(date.Date)
				if err != nil {
					return nil, err
				}
				if timestamp.Before(t) {
					newComic := Comic{Name: c.Title, Date: date.Date, Prices: c.Prices, Urls: c.Urls}
					publishedComics = append(publishedComics, newComic)
					break
				}
			}
		}
	}
	return publishedComics, nil
}
