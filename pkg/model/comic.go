package model

import "time"

type Comic struct {
	Name   string
	Date   time.Time
	Prices []Price
	Urls   []BuyUrl
}

type Price struct {
	PriceType string
	Value     float64
}

type BuyUrl struct {
	UrlType string
	Url     string
}
