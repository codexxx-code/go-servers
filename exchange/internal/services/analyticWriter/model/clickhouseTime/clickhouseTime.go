package clickhouseTime

import (
	"fmt"
	"time"
)

type ClickhouseTime struct {
	time.Time
}

const layout = "2006-01-02 15:04:05"

func (ct *ClickhouseTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(layout, string(b))
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct ClickhouseTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%v"`, ct.Format(layout))), nil
}

func Now() ClickhouseTime {
	return ClickhouseTime{Time: time.Now()}
}
