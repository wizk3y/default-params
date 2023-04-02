package internal

import (
	"reflect"
	"strconv"
)

// CastValueTo -- convert value from string to specific reflect.Type
func CastValueTo(strValue string, vType reflect.Type, ptr bool) interface{} {
	if assert := getAssert(vType); assert != nil {
		if ptr {
			return assert.ToPtrType(strValue)
		}

		return assert.ToType(strValue)
	}

	switch vType.Kind() {
	case reflect.String:
		if ptr {
			return StringPtr(strValue)
		}

		return strValue
	case reflect.Bool:
		defaultValue, _ := strconv.ParseBool(strValue)

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
		defaultValue, _ := strconv.ParseInt(strValue, 10, 64)

		if ptr {
			return Int64Ptr(defaultValue)
		}

		return defaultValue
	default:
		return nil
	}
}
