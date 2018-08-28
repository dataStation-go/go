package config

import (
	"strings"
	"strconv"
	"github.com/mytokenio/go_sdk/log"
)

func toBool(v interface{}, dv bool) bool {
	s := strings.ToLower(toString(v, ""))
	if s == "true" {
		return true
	} else if s == "false" {
		return false
	}

	if x, ok := v.(bool); ok {
		return x
	}
	return dv
}

func toString(v interface{}, dv string) string {
	if x, ok := v.(string); ok {
		return x
	}
	return dv
}

func toInt(v interface{}, dv int) int {
	switch x := v.(type) {
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			log.Errorf("failed convert %s to int64 %s", x, err)
			return dv
		}
		return int(i)
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	}
	return dv
}

func toInt64(v interface{}, dv int64) int64 {
	switch x := v.(type) {
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			log.Errorf("failed convert %s to int64 %s", x, err)
			return dv
		}
		return i
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	case float32:
		return int64(x)
	case float64:
		return int64(x)
	}
	return dv
}

func toFloat64(v interface{}, dv float64) float64 {
	switch x := v.(type) {
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			log.Errorf("failed convert %s to float64 %s", x, err)
			return dv
		}
		return float64(f)
	case float32:
		return float64(x)
	case float64:
		return x
	case int:
		return float64(x)
	case int8:
		return float64(x)
	case int16:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)
	}
	return dv
}