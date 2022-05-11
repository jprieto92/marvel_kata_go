package comics_test

import (
	"encoding/json"
	"fmt"
	"github.com/jprieto92/marvel_kata_go/src/pkg/comics"
	"github.com/jprieto92/marvel_kata_go/src/pkg/model"
	"github.com/jprieto92/marvel_kata_go/src/pkg/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetComics(t *testing.T) {
	t.Run("malformed uri", func(t *testing.T) {
		uri := "golang.org"
		_, err := comics.NewComicProcessor(uri)

		assertError(t, err)
	})
	t.Run("comics server returns no ok", func(t *testing.T) {
		server := makeServerNok()
		defer server.Close()

		processor, err := comics.NewComicProcessor(server.URL)
		assertNonError(t, err)
		_, err = processor.GetAllComics()
		assertError(t, err)

	})
	t.Run("return all comics", func(t *testing.T) {
		server := makeServer()
		defer server.Close()

		processor, err := comics.NewComicProcessor(server.URL)
		assertNonError(t, err)
		got, err := processor.GetAllComics()

		want, _ := os.ReadFile("./testdata/comicsdb_response.json")
		assertNonError(t, err)
		assertResponse(t, got, string(want))
	})
}

func TestGetComicsPublishedThisWeekUntilNow(t *testing.T) {
	t.Run("return comics from this week until now", func(t *testing.T) {
		server := makeServer()
		defer server.Close()

		processor, err := comics.NewComicProcessor(server.URL)
		assertNonError(t, err)
		timeNow, _ := utils.ConvertTimestampToDate("2022-05-11T00:00:00-0400")
		got, err := processor.GetComicsPublishedInWeekUntilTime(timeNow)

		want, _ := json.Marshal([]model.Comic{{Name: "Spider-Verse Unlimited Infinity Comic (2022) #5",
			Date:   "2022-05-10T00:00:00-0400",
			Prices: []model.Price{{Type: "printPrice", Price: 0.0}},
			Urls:   []model.Url{{Type: "detail", URL: "http://marvel.com/comics/issue/98596/spider-verse_unlimited_infinity_comic_2022_5?utm_campaign=apiRef&utm_source=97f295907072a970c5df30d73d1f3816"}}},
			{Name: "X-Men Unlimited Infinity Comic (2021) #34",
				Date:   "2022-05-09T00:00:00-0400",
				Prices: []model.Price{{Type: "printPrice", Price: 0.0}},
				Urls:   []model.Url{{Type: "detail", URL: "http://marvel.com/comics/issue/101322/x-men_unlimited_infinity_comic_2021_34?utm_campaign=apiRef&utm_source=97f295907072a970c5df30d73d1f3816"}}},
		})
		assertNonError(t, err)
		assertResponse(t, got, string(want))
	})
	t.Run("server fails when try to get comics published this week ", func(t *testing.T) {
		server := makeServerNok()
		defer server.Close()

		processor, err := comics.NewComicProcessor(server.URL)
		assertNonError(t, err)
		timeNow, _ := utils.ConvertTimestampToDate("2022-05-11T00:00:00-0400")
		_, err = processor.GetComicsPublishedInWeekUntilTime(timeNow)
		assertError(t, err)
	})
	t.Run("server retrieves malformed json when try to get comics published this week ", func(t *testing.T) {
		server := makeServerMalformedJson()
		defer server.Close()

		processor, err := comics.NewComicProcessor(server.URL)
		assertNonError(t, err)
		timeNow, _ := utils.ConvertTimestampToDate("2022-05-11T00:00:00-0400")
		_, err = processor.GetComicsPublishedInWeekUntilTime(timeNow)
		assertError(t, err)
	})
}

func assertResponse(t *testing.T, got, want string) {
	if got != string(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, got error) {
	t.Helper()
	if got == nil {
		t.Errorf("didn't get an error but wanted one")
	}
}

func assertNonError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Get an error but not wanted")
	}
}

func makeServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		file, _ := os.ReadFile("./testdata/comicsdb_response.json")
		fmt.Fprintf(w, string(file))
	}))
}

func makeServerNok() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
}

func makeServerMalformedJson() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "---")
	}))
}
