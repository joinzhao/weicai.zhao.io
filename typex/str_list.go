package typex

import (
	"bytes"
	"errors"
	"strings"
)

const (
	stringSep = ","
)

type StrList []string

func (v StrList) MarshalJSON() ([]byte, error) {
	if len(v) == 0 {
		return []byte(""), nil
	}

	str := strings.Join(v, stringSep)

	return []byte(str), nil
}

func (v *StrList) UnmarshalJSON(data []byte) error {
	if data == nil {
		return errors.New("null point exception")
	}
	if bytes.Equal(data, emptyString) {
		return nil
	}

	list := strings.Split(string(data), stringSep)

	*v = append((*v)[0:0], list...)

	return nil
}
