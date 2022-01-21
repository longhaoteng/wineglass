package cron

type Option func(*Options)

type Options struct {
	Wrapper    *JobWrapper
	SingleNode bool
}

func Wrapper(wrapper *JobWrapper) Option {
	return func(o *Options) {
		o.Wrapper = wrapper
	}
}

func SingleNode(singleNode bool) Option {
	return func(o *Options) {
		o.SingleNode = singleNode
	}
}
