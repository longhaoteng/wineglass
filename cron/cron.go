package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/lock"
)

var (
	jobs []Job
	c    *cron.Cron

	EmptyOptions []Option

	Recover             = cron.Recover(newLog())
	SkipIfStillRunning  = cron.SkipIfStillRunning(newLog())
	DelayIfStillRunning = cron.DelayIfStillRunning(newLog())
)

type JobWrapper = cron.JobWrapper

type Job interface {
	Name() string
	Spec() string
	Options() []Option
	Run()
}

func Init() error {
	shanghai, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	c = cron.New(cron.WithSeconds(), cron.WithLocation(shanghai))

	for _, job := range jobs {
		opt := &Options{}
		for _, o := range job.Options() {
			o(opt)
		}

		var cronJob cron.Job = job

		if opt.SingleNode {
			// loop variables captured by 'func' literals in 'go' statements might have unexpected values
			entryJob := job
			cronJob = cron.FuncJob(func() {
				l := lock.NewLock(fmt.Sprintf("cron_single_node_%s_%s", config.Service.Name, entryJob.Name()))
				if !l.Lock() {
					return
				}
				entryJob.Run()
				l.UnLock()
			})
		}

		if len(opt.Wrappers) > 0 {
			cronJob = cron.NewChain(opt.Wrappers...).Then(cronJob)
		}

		_, err = c.AddJob(job.Spec(), cronJob)
		if err != nil {
			return err
		}
	}

	c.Start()

	return nil
}

func AddJobs(jos ...Job) {
	jobs = append(jobs, jos...)
}

func Stop() {
	c.Stop()
}
