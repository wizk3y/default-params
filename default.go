package defaults

import (
	"errors"
	"reflect"
	"strings"

	"github.com/wizk3y/default-params/internal"
)

var (
	extractTag string = "default"
	nilSign    string = "-"
)

// SetExtractTag --
func SetExtractTag(tag string) {
	extractTag = tag
}

// SetNilSign --
func SetNilSign(sign string) {
	nilSign = sign
}

// MustFillDefaultValue --
func MustFillDefaultValue(i interface{}, opt ...bool) {
	err := FillDefaultValue(i, opt...)
	if err != nil {
		panic(err)
	}
}

// FillDefaultValue --
func FillDefaultValue(i interface{}, opt ...bool) error {
	skipNonZero := true
	if len(opt) > 0 {
		skipNonZero = opt[0]
	}

	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr {
		return errors.New("input should be pointer to fill value")
	}

	if v.IsNil() {
		return errors.New("nil pointer cant set value")
	}

	ve := v.Elem()
	t := reflect.TypeOf(i).Elem()
	for index := 0; index < t.NumField(); index++ {
		f := t.Field(index)

		if !ve.FieldByName(f.Name).IsZero() && skipNonZero {
			continue
		}

		setValueByType(ve, f.Name, f.Type, f.Tag.Get(extractTag))
	}

	return nil
}

func setValueByType(ve reflect.Value, name string, vType reflect.Type, valueStr string, opt ...bool) {
	ptr := false
	if len(opt) > 0 {
		ptr = opt[0]
	}

	switch vType.Kind() {
	case reflect.String, reflect.Bool, reflect.Float64, reflect.Float32, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value := internal.CastValueTo(valueStr, vType, ptr)

		ve.FieldByName(name).Set(reflect.ValueOf(value))
	case reflect.Struct:
		defaultValue := reflect.New(vType)

		FillDefaultValue(defaultValue.Interface())

		if ptr {
			ve.FieldByName(name).Set(defaultValue)
		} else {
			ve.FieldByName(name).Set(defaultValue.Elem())
		}
	case reflect.Ptr:
		if valueStr == nilSign {
			return
		}

		setValueByType(ve, name, internal.Deref(vType), valueStr, true)
	case reflect.Slice:
		defaultValues := strings.Split(valueStr, ",")

		dataValue := reflect.MakeSlice(vType, 0, 0)

		for _, vStr := range defaultValues {
			var value interface{}

			switch vType.Elem().Kind() {
			case reflect.String, reflect.Bool, reflect.Float64, reflect.Float32, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value = internal.CastValueTo(vStr, vType.Elem())
			case reflect.Ptr:
				value = internal.CastValueTo(vStr, internal.Deref(vType.Elem()), true)
			}

			dataValue = reflect.Append(dataValue, reflect.ValueOf(value))
		}

		ve.FieldByName(name).Set(dataValue)
	}
}
