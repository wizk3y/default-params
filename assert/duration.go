package assert

import (
	"fmt"
	"strings"
	"time"

	"github.com/wizk3y/default-params/internal"
)

type durationTypeAssert struct {
	intSuffix string
}

// NewDurationTypeAssert --
func NewDurationTypeAssert(intSuffix string) internal.TypeAssert {
	return &durationTypeAssert{
		intSuffix: intSuffix,
	}
}

func (a *durationTypeAssert) ToType(strValue string) interface{} {
	if !containsDurationSuffix(strValue) {
		strValue = fmt.Sprintf("%s%s", strValue, a.intSuffix)
	}

	defaultValue, _ := time.ParseDuration(strValue)

	return defaultValue
}

func (a *durationTypeAssert) ToPtrType(strValue string) interface{} {
	defaultValue := a.ToType(strValue).(time.Duration)

	return internal.DurationPtr(defaultValue)
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
