package tools

import (
	"bytes"
	"encoding/json"
	"io"
)

// AnyReader any to io.Reader
func AnyReader(v any) (io.Reader, error) {
	switch v.(type) {
	case io.Reader:
		return v.(io.Reader), nil
	}

	return anyReader(v)
}

func anyReader(v any) (io.Reader, error) {
	var (
		data, err = json.Marshal(v)
	)
	if err != nil {
		return nil, nil
	}
	return bytes.NewReader(data), nil
}
