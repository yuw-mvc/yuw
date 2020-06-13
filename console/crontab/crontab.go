package crontab

import (
	"github.com/robfig/cron"
	M "github.com/yuw-mvc/yuw/modules"
)

func Do(cronTabPoT *PoT) {
	if M.I.Get("YuwConsole.Crontab", false).(bool) {
		crons = cronTabPoT.Crons
		if crons != nil {
			go new(cronTab).do()
		}
	}
}

var crons *PoTCronTabs

type (
	PoTCronTabs map[string][]CronJob
	CronJob cron.Job

	cronTab struct {

	}

	PoT struct {
		Crons *PoTCronTabs
	}
)

func (cTab *cronTab) do() {
	c := cron.New()

	for spec, cmds := range *crons {
		if cmds == nil {
			continue
		}

		s, err := cron.Parse(spec)
		if err != nil {
			continue
		}

		for _, cmd := range cmds {
			c.Schedule(s, cmd)
		}
	}

	c.Start()
	defer c.Stop()

	select {}
}
