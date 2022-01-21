package jobs

import (
	"github.com/longhaoteng/wineglass/cron"
	log "github.com/longhaoteng/wineglass/logger"
)

type Test struct{}

func (t *Test) Name() string {
	return "test"
}

func (t *Test) Spec() string {
	return "@every 5m"
}

func (t *Test) Options() []cron.Option {
	return []cron.Option{
		cron.Wrapper(&cron.DelayIfStillRunning),
	}
}

func (t *Test) Run() {
	t.execute()
}

func (t *Test) execute() {
	log.Info("coon test")
}
