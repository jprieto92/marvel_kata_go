package utils

import (
	"fmt"
	"time"
)

var ErrDecodingTimestamp = string("error when try to decode timestamp")

//ConvertTimestampToDate convert time from ISO8661 to time.Time()
func ConvertTimestampToDate(timestamp string) (time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04:05-0700", timestamp)
	if err != nil {
		return time.Time{}, fmt.Errorf("%v: %w", ErrDecodingTimestamp, err)
	}

	return t, nil
}
