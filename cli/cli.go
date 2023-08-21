package cli

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	fileFlag      = "configfile"
	fileShortFlag = "f"
)

func RunServiceCli() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     fileFlag,
				Aliases:  []string{fileShortFlag},
				Usage:    "file path",
				Required: true,
			},
		},
		Action: StartService,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to start service %v", err)
		return
	}
}

func StartService(c *cli.Context) error {

	if c.String(fileFlag) == "" {
		return errors.New("param error")
	}

	return nil
}
