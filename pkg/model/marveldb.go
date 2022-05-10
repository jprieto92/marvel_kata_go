package model

type MarvelDb struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Count   int `json:"count"`
		Results []struct {
			ID                 int    `json:"id"`
			DigitalID          int    `json:"digitalId"`
			Title              string `json:"title"`
			IssueNumber        int    `json:"issueNumber"`
			VariantDescription string `json:"variantDescription"`
			Description        string `json:"description"`
			Modified           string `json:"modified"`
			Isbn               string `json:"isbn"`
			Upc                string `json:"upc"`
			DiamondCode        string `json:"diamondCode"`
			Ean                string `json:"ean"`
			Issn               string `json:"issn"`
			Format             string `json:"format"`
			PageCount          int    `json:"pageCount"`
			TextObjects        []struct {
				Type     string `json:"type"`
				Language string `json:"language"`
				Text     string `json:"text"`
			} `json:"textObjects"`
			ResourceURI string `json:"resourceURI"`
			Urls        []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"urls"`
			Series struct {
				ResourceURI string `json:"resourceURI"`
				Name        string `json:"name"`
			} `json:"series"`
			Variants        []interface{} `json:"variants"`
			Collections     []interface{} `json:"collections"`
			CollectedIssues []interface{} `json:"collectedIssues"`
			Dates           []struct {
				Type string `json:"type"`
				Date string `json:"date"`
			} `json:"dates"`
			Prices []struct {
				Type  string  `json:"type"`
				Price float64 `json:"price"`
			} `json:"prices"`
			Thumbnail struct {
				Path      string `json:"path"`
				Extension string `json:"extension"`
			} `json:"thumbnail"`
			Images []struct {
				Path      string `json:"path"`
				Extension string `json:"extension"`
			} `json:"images"`
			Creators struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
					Role        string `json:"role"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"creators"`
			Characters struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"characters"`
			Stories struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"stories"`
			Events struct {
				Available     int           `json:"available"`
				CollectionURI string        `json:"collectionURI"`
				Items         []interface{} `json:"items"`
				Returned      int           `json:"returned"`
			} `json:"events"`
		} `json:"results"`
	} `json:"data"`
}
