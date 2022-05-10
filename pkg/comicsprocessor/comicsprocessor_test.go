package comicsprocessor_test

import (
	"fmt"
	"github.com/jprieto92/marvel_kata_go/pkg/comicsprocessor"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetComics(t *testing.T) {
	t.Run("malformed uri", func(t *testing.T) {
		uri := "golang.org"
		_, err := comicsprocessor.NewComicProcessor(uri)

		assertError(t, err, comicsprocessor.ErrBadUri)
	})
	t.Run("comics server returns no ok", func(t *testing.T) {
		server := makeServerNok()
		defer server.Close()

		processor, err := comicsprocessor.NewComicProcessor(server.URL)
		assertNonError(t, err)
		_, err = processor.GetAllComics()
		assertError(t, err, comicsprocessor.ErrStatusCodeNoOk)

	})
	t.Run("return all comics", func(t *testing.T) {
		server := makeServer()
		defer server.Close()

		processor, err := comicsprocessor.NewComicProcessor(server.URL)
		assertNonError(t, err)
		got, err := processor.GetAllComics()

		want, _ := os.ReadFile("./testdata/comicsdb_response.json")
		assertNonError(t, err)
		assertResponse(t, got, string(want))
	})
}

func assertResponse(t *testing.T, got, want string) {
	if got != string(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
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
