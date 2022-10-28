package cron

type Option func(*Options)

type Options struct {
	Wrappers   []JobWrapper
	SingleNode bool
}

func Wrappers(wrappers ...JobWrapper) Option {
	return func(o *Options) {
		o.Wrappers = append(o.Wrappers, wrappers...)
	}
}

func SingleNode() Option {
	return func(o *Options) {
		o.SingleNode = true
	}
}
