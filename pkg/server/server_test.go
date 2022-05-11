package server_test

import (
	"fmt"
	"github.com/jprieto92/marvel_kata_go/pkg/server"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type FakeDbInfo struct {
	DbUri string
}

func (db *FakeDbInfo) GetDatabaseUri() string {
	return db.DbUri
}

func TestHandleRequestHandler(t *testing.T) {
	t.Run("succesfull response", func(t *testing.T) {
		fakeDbServer := makeServer()
		defer fakeDbServer.Close()

		dbInfo := FakeDbInfo{DbUri: fakeDbServer.URL}
		serv := &server.Server{&dbInfo}

		req := httptest.NewRequest(http.MethodGet, "/comics", nil)
		w := httptest.NewRecorder()
		serv.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("internal server error response", func(t *testing.T) {
		fakeDbServer := makeEmptyServer()
		defer fakeDbServer.Close()

		dbInfo := FakeDbInfo{DbUri: fakeDbServer.URL}
		serv := &server.Server{&dbInfo}

		req := httptest.NewRequest(http.MethodGet, "/comics", nil)
		w := httptest.NewRecorder()
		serv.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusInternalServerError

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("method not allowed response", func(t *testing.T) {
		fakeDbServer := makeEmptyServer()
		defer fakeDbServer.Close()

		dbInfo := FakeDbInfo{DbUri: fakeDbServer.URL}
		serv := &server.Server{&dbInfo}

		req := httptest.NewRequest(http.MethodPost, "/comics", nil)
		w := httptest.NewRecorder()
		serv.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusMethodNotAllowed

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func makeServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		file, _ := os.ReadFile("./testdata/comicsdb_response.json")
		fmt.Fprintf(w, string(file))
	}))
}

func makeEmptyServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "[]")
	}))
}
