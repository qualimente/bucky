package main

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "bucky"
	app.Usage = "Produce a periodic bucky message"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "verbose",
			Usage: "Print verbose output",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print the version and exit",
			Action: func(c *cli.Context) error {
				fmt.Printf("bucky version %s\n", app.Version)
				return nil
			},
		},
		{
			Name:   "generate",
			Usage:  "Generate log messages",
			Action: RunPeriodically,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "format",
					Usage: "Specify the output format: text, json",
				}, cli.DurationFlag{
					Name:  "period",
					Usage: "Specify the period between starting log message emitters. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' (default: 1s)",
					Value: 1 * time.Second,
				}, cli.Uint64Flag{
					Name:  "count",
					Usage: "Specify the count of messages emitted per period",
					Value: 1,
				},
			},
		},
	}

	app.RunAndExitOnError()
}
