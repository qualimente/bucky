package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
	"math/rand"
	"github.com/icrowley/fake"
	"fmt"
)

func RunPeriodically(c *cli.Context) error {
	rand.Seed(time.Now().UTC().UnixNano())

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
	tNanos := time.Now().UTC().UnixNano()
	for i := uint64(0); i < messageCount; i++ {
		msgId := fmt.Sprintf("%d-%d", tNanos, i)
		log.WithFields(log.Fields{
			"type": "log",
			"msgId": msgId,
		}).Info(fake.SentencesN(rand.Intn(10)))
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
