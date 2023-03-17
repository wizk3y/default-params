package defaults_test

import (
	"fmt"
	"reflect"
	"testing"

	defaults "github.com/wizk3y/default-params"
)

// custom type `Sample` from primitive type `string`
type Sample string

func samplePtr(s Sample) *Sample {
	return &s
}

// implement struct with 2 func `ToType` and `ToPtrType`
type sampleTypeAssert struct{}

func (a *sampleTypeAssert) ToType(strValue string) interface{} {
	return Sample(fmt.Sprintf("sample-%s", strValue))
}

func (a *sampleTypeAssert) ToPtrType(strValue string) interface{} {
	return samplePtr(a.ToType(strValue).(Sample))
}

// ---
type AssertParameter struct {
	Sample    Sample  `default:"text123" json:"sample"`
	SamplePtr *Sample `default:"text123" json:"sample_ptr"`
}

func Test_RegisterAssertFunc(t *testing.T) {
	typeAssert := sampleTypeAssert{}
	defaults.RegisterAssertFunc(reflect.String, "github.com/wizk3y/default-params_test.Sample", &typeAssert)

	params := AssertParameter{}

	defaults.MustFillDefaultValue(&params)

	equal := reflect.DeepEqual(params, AssertParameter{
		Sample:    Sample("sample-text123"),
		SamplePtr: samplePtr(Sample("sample-text123")),
	})
	if !equal {
		t.Fatal("output not equal")
	}
}
