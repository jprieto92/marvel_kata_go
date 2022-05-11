package model

import (
	"encoding/json"
	"fmt"
)

const ErrMarshallingJson = "error when try marshalling comics"

type ListComics []Comic

type Comic struct {
	Name   string  `json:"title"`
	Date   string  `json:"date"`
	Prices []Price `json:"prices"`
	Urls   []Url   `json:"urls"`
}

func (c *ListComics) EncodeComicListToJson() (string, error) {
	if len(*c) > 0 {
		result, err := json.Marshal(c)
		if err != nil {
			return "", fmt.Errorf("%v: %w", ErrMarshallingJson, err)
		}
		return string(result), nil
	}

	return "[]", nil
}
