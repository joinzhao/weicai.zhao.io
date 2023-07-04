package typex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guregu/null"
	"time"
)

var nullBytes = []byte("null")
var emptyString = []byte("\"\"")

type lt = null.Time
type LocalTime struct {
	lt
}

// MarshalJSON 重写 null.Time 的序列化方法
func (t LocalTime) MarshalJSON() ([]byte, error) {

	if !t.Valid {
		return []byte("null"), nil
	}

	tune := t.Time.In(time.Local).Format(`"2006-01-02 15:04:05"`)

	return []byte(tune), nil
}

// UnmarshalJSON 重写null.time 反序列化方法
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) || bytes.Equal(data, emptyString) {
		t.Valid = false
		return nil
	}

	var err error
	flag := false

	layouts := []string{"2006-01-02 15:04:05", "2006-01-02 15:04", "2006-01-02", "", "2006", "2006-01",
		"2006/01/02 15:04:05", "2006/01/02 15:04", "2006/01/02", "2006/01"}

	for _, layout := range layouts {
		var s string

		// to string
		if err := json.Unmarshal(data, &s); err != nil {
			continue
		}

		var timestamp time.Time

		// to timestamp
		if timestamp, err = time.ParseInLocation(layout, s, time.Local); err != nil {
			continue
		}

		var bts []byte

		// to []byte
		if bts, err = timestamp.MarshalJSON(); err != nil {
			continue
		}

		if err = json.Unmarshal(bts, &t.Time); err != nil {
			continue
		}

		flag = true

		break
	}

	if !flag {
		return fmt.Errorf("null: couldn't unmarshal JSON: %w", err)
	}

	t.Valid = true
	return nil

}
