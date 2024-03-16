package Players

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = strings.Trim(str, "\"")
	parsed, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	ct.Time = parsed
	return nil
}
