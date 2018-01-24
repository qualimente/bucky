package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
)

func RunPeriodically(c *cli.Context) error {

	log.SetFormatter(_makeFormatter(c.String("format")))

	log.WithFields(log.Fields{
		"appName": c.App.Name,
	}).Info("Running periodically")

	period := c.Duration("period")

	messageCount := c.Uint64("count")

	for {
		go func() {
			PrintMessages(messageCount)
		}()

		time.Sleep(period)
	}

}

func PrintMessages(messageCount uint64) {
	for i := uint64(0); i < messageCount; i++ {
		log.WithFields(log.Fields{
			"type": "log",
		}).Info("Every heartbeat bears your name")
	}
}

func _makeFormatter(format string) log.Formatter {
	switch format {
	case "text":
		return &log.TextFormatter{DisableColors: true}
	case "json":
		return &log.JSONFormatter{}
	default:
		return &log.JSONFormatter{}
	}
}
