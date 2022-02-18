package common

import (
	"encoding/json"
	"strings"
	"time"
)

func (j *VOTODate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = VOTODate(t)
	return nil
}

func (j VOTODate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j VOTODate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j VOTODate) Time() *time.Time {
	t := time.Time(j)
	if !t.IsZero() {
		return &t
	}
	return nil
}
