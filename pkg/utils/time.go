package utils

import (
	"fmt"
	"time"
)

const ErrDecodingTimestamp = "error when try to decode timestamp"
const TimeLayout = "2006-01-02T15:04:05-0700"

//ConvertTimestampToDate convert time from ISO8661 to time.Time()
func ConvertTimestampToDate(timestamp string) (time.Time, error) {
	t, err := time.Parse(TimeLayout, timestamp)
	if err != nil {
		return time.Time{}, fmt.Errorf("%v: %w", ErrDecodingTimestamp, err)
	}

	return t, nil
}
