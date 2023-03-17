package defaults

import "github.com/wizk3y/default-params/internal"

type overrideSettedValueOpt bool

// OverrideSettedValueOpt --
func OverrideSettedValueOpt() internal.FillOpt {
	return overrideSettedValueOpt(true)
}

func (o overrideSettedValueOpt) Apply(conf *internal.FillConfig) {
	conf.OverrideSettedValue = bool(o)
}
