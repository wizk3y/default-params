package internal

import (
	"reflect"
	"time"
)

// BoolPtr --
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtr --
func StringPtr(s string) *string {
	return &s
}

// IntPtr --
func IntPtr(i int) *int {
	return &i
}

// Int8Ptr --
func Int8Ptr(i int8) *int8 {
	return &i
}

// Int16Ptr --
func Int16Ptr(i int16) *int16 {
	return &i
}

// Int3Ptr --
func Int3Ptr(i int32) *int32 {
	return &i
}

// Int64Ptr --
func Int64Ptr(i int64) *int64 {
	return &i
}

// Float32Ptr --
func Float32Ptr(f float32) *float32 {
	return &f
}

// Float64Ptr --
func Float64Ptr(f float64) *float64 {
	return &f
}

// DurationPtr --
func DurationPtr(d time.Duration) *time.Duration {
	return &d
}

// Deref --
func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}
