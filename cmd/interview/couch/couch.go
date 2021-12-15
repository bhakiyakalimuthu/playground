package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func convert(data interface{}) (string, error) {
	if data == nil {
		return "", errors.New("data is nil")
	}

	switch v := data.(type) {
	case string:
		if v == "" {
			return "", errors.New("data is empty string")
		}

		v = strings.ReplaceAll(v, "$", "")
		v = strings.ReplaceAll(v, ",", "_")
		s, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return "", err
		}
		return strconv.FormatFloat(s, 'f', 2, 64), nil
	case int:
		return strconv.FormatFloat(float64(v), 'f', 2, 64), nil
	case int32:
		return strconv.FormatFloat(float64(v), 'f', 2, 32), nil
	case int64:
		return strconv.FormatFloat(float64(v), 'f', 2, 64), nil
	case float32:
		// return strconv.FormatFloat(float64(v), 'F', 2, 64), nil
		return fmt.Sprintf("%f", v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', 2, 32), nil
	case []byte:
		return convert(string(v))
	default:
		return "", errors.New("unknown interface")
	}
}
