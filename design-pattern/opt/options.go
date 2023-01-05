package opt

import "fmt"

type Options struct {
	A string
	B string
}

type Opt func(*Options)

func InitOptions(opts ...Opt) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	fmt.Printf("%+v", opt)
}

func WithOptionA(a string) Opt {
	return func(o *Options) {
		o.A = a
	}
}

func WithOptionB(b string) Opt {
	return func(o *Options) {
		o.B = b
	}
}
