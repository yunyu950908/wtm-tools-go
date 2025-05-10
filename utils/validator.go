package utils

import (
	"github.com/robfig/cron/v3"
)

var CronParser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

func ValidateCronExpr(expr string) error {
	_, err := CronParser.Parse(expr)
	return err
}
