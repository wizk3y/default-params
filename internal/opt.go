package internal

type FillOpt interface {
	Apply(conf *FillConfig)
}
