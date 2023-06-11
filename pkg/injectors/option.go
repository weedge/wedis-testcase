package injectors

type OptPrinter interface {
	String() string
}

type Option interface {
	apply(OptPrinter)
}

type Opt struct {
	f func(OptPrinter)
}

func (fdo *Opt) apply(do OptPrinter) {
	fdo.f(do)
}

func NewOpt(f func(OptPrinter)) *Opt {
	return &Opt{
		f: f,
	}
}
