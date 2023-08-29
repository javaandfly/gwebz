package gwebz

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

const (
	fileFlag      = "configfile"
	fileShortFlag = "f"
)

func RunSyncServerCli() {
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
		return
	}
}

func StartService(c *cli.Context) error {
	if c.String(fileFlag) == "" {
		return errors.New("no config file specified")
	}

	r := gin.Default()
	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}
