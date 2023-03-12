package internal

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

// CastValueTo -- convert value from string to specific reflect.Type
func CastValueTo(strValue string, vType reflect.Type, opt ...bool) interface{} {
	ptr := false
	if len(opt) > 0 {
		ptr = opt[0]
	}

	switch vType.Kind() {
	case reflect.String:
		if ptr {
			return StringPtr(strValue)
		}

		return strValue
	case reflect.Bool:
		defaultValue := false
		if strValue == "true" || strValue == "1" {
			defaultValue = true
		}

		if ptr {
			return BoolPtr(defaultValue)
		}

		return defaultValue
	case reflect.Float64:
		defaultValue, _ := strconv.ParseFloat(strValue, 64)

		if ptr {
			return Float64Ptr(defaultValue)
		}

		return defaultValue
	case reflect.Float32:
		defaultValue, _ := strconv.ParseFloat(strValue, 32)

		if ptr {
			return Float64Ptr(defaultValue)
		}

		return defaultValue
	case reflect.Int:
		defaultValue, _ := strconv.ParseInt(strValue, 10, 0)

		if ptr {
			return Int64Ptr(defaultValue)
		}

		return defaultValue
	case reflect.Int8:
		defaultValue, _ := strconv.ParseInt(strValue, 10, 8)

		if ptr {
			return Int64Ptr(defaultValue)
		}

		return defaultValue
	case reflect.Int16:
		defaultValue, _ := strconv.ParseInt(strValue, 10, 16)

		if ptr {
			return Int64Ptr(defaultValue)
		}

		return defaultValue
	case reflect.Int32:
		defaultValue, _ := strconv.ParseInt(strValue, 10, 32)

		if ptr {
			return Int64Ptr(defaultValue)
		}

		return defaultValue
	case reflect.Int64:
		if vType.String() == "time.Duration" {
			if !containsDurationSuffix(strValue) {
				strValue += "ns"
			}

			defaultValue, _ := time.ParseDuration(strValue)

			if ptr {
				return DurationPtr(defaultValue)
			}

			return defaultValue
		}

		defaultValue, _ := strconv.ParseInt(strValue, 10, 64)

		if ptr {
			return Int64Ptr(defaultValue)
		}

		return defaultValue
	default:
		return nil
	}
}

func containsDurationSuffix(s string) bool {
	validSuffixes := []string{"h", "m", "s", "ms", "µs", "us", "ns"}

	for _, vs := range validSuffixes {
		if strings.HasSuffix(s, vs) {
			return true
		}
	}

	return false
}