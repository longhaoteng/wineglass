package retry

import (
	"time"

	rty "github.com/avast/retry-go"
)

var (
	Delay     = rty.Delay
	Attempts  = rty.Attempts
	MaxDelay  = rty.MaxDelay
	MaxJitter = rty.MaxJitter
	OnRetry   = rty.OnRetry

	DefaultRetry = NewRetry("default", Delay(100*time.Millisecond), Attempts(uint(5)))
)

type (
	Option        = rty.Option
	RetryableFunc = rty.RetryableFunc
)

type Retry struct {
	name string
	opts []Option
}

func NewRetry(name string, opts ...Option) *Retry {
	return &Retry{
		name: name,
		opts: opts,
	}
}

func (r *Retry) Name() string {
	return r.name
}

func (r *Retry) Do(retryableFunc RetryableFunc, opts ...Option) error {
	return rty.Do(retryableFunc, append(r.opts, opts...)...)
}
