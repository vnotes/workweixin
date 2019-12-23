package schedules

import (
	"github.com/robfig/cron/v3"
	"github.com/vnotes/workweixin_app/cores/notifications"
)

func RegisterCronJob(c *cron.Cron) {
	_, _ = c.AddFunc("0 1 * * *", func() {
		notifications.AppMsgPush()
	})
	c.Start()
}