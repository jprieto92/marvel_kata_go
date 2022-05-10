package utils_test

import (
	"github.com/jprieto92/marvel_kata_go/pkg/utils"
	"testing"
)

func TestConvertTimestampToDate(t *testing.T) {
	t.Run("wrong timestamp", func(t *testing.T) {
		timestamp := "2022-05-10X00:00:00-0400"
		_, err := utils.ConvertTimestampToDate(timestamp)

		assertError(t, err)
	})
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
