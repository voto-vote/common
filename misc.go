package common

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

var PW_LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!§$%&/()=?`*'_:;<>,.-")

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

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = PW_LETTERS[rand.Intn(len(PW_LETTERS))]
	}
	return string(b)
}

func GetAge(birthdate time.Time) int {
	today := time.Now()
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}
