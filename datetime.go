package tilda_go

import (
	"strings"
	"time"
)

// DateTime is a custom type for time.Time that allows to unmarshal JSON with a specific format
type DateTime time.Time

const dtFormat = "2006-01-02 15:04:05"

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*d = DateTime(time.Time{})
		return nil
	}
	t, err := time.Parse(dtFormat, s)
	if err != nil {
		return err
	}

	*d = DateTime(t)
	return nil
}
