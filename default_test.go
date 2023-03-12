package defaults_test

import (
	"reflect"
	"testing"
	"time"

	defaults "github.com/wizk3y/default-params"
	"github.com/wizk3y/default-params/internal"
)

type AbstractParams struct {
	AbsBool  bool    `default:"true" json:"abs_bool"`
	AbsFloat float64 `default:"2.2" json:"abs_float"`
}

type SubParams struct {
	Str string `default:"abc" json:"str"`
	Int int64  `default:"2" json:"int"`
}

type Parameters struct {
	Str         string     `default:"abc" json:"str"`
	Int         int64      `default:"2" json:"int"`
	Bool        bool       `default:"true" json:"bool"`
	Float       float64    `default:"2.2" json:"float"`
	StrPtr      *string    `default:"def" json:"str_ptr"`
	StrSetted   string     `default:"abc" json:"str_setted"`
	IntPtr      *int64     `default:"22" json:"int_ptr"`
	BoolPtr     *bool      `default:"true" json:"bool_ptr"`
	FloatPtr    *float64   `default:"3.3" json:"float_ptr"`
	SliceStr    []string   `default:"1,2,3" json:"slice_str"`
	SliceStrPtr []*string  `default:"a,b,c" json:"slice_str_ptr"`
	Sub         SubParams  `json:"sub"`
	SubPtr      *SubParams `json:"sub_ptr"`
	*AbstractParams
	StrDur    time.Duration `default:"10ms" json:"str_dur"`
	IntDurPtr time.Duration `default:"1000000000" json:"int_dur_ptr"`
}

func Test_MustFillDefaultValue(t *testing.T) {
	params := Parameters{
		StrSetted: "setted",
	}

	defaults.MustFillDefaultValue(&params)

	equal := reflect.DeepEqual(params, Parameters{
		Str:         "abc",
		Int:         2,
		Bool:        true,
		Float:       2.2,
		StrPtr:      internal.StringPtr("def"),
		StrSetted:   "setted",
		IntPtr:      internal.Int64Ptr(22),
		BoolPtr:     internal.BoolPtr(true),
		FloatPtr:    internal.Float64Ptr(3.3),
		SliceStr:    []string{"1", "2", "3"},
		SliceStrPtr: []*string{internal.StringPtr("a"), internal.StringPtr("b"), internal.StringPtr("c")},
		Sub: SubParams{
			Str: "abc",
			Int: 2,
		},
		SubPtr: &SubParams{
			Str: "abc",
			Int: 2,
		},
		AbstractParams: &AbstractParams{
			AbsBool:  true,
			AbsFloat: 2.2,
		},
		StrDur:    time.Duration(10) * time.Millisecond,
		IntDurPtr: *internal.DurationPtr(time.Duration(1) * time.Second),
	})
	if equal {
		t.Error("output not equal")
		return
	}

	// p, err := json.Marshal(&params)
	// fmt.Println(string(p))
	// fmt.Println(err)
}

func Benchmark_MustFillDefaultValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		defaults.MustFillDefaultValue(&Parameters{})
	}
}
