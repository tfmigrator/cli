package cli

import (
	"fmt"

	"github.com/tfmigrator/cli/pkg/controller"
	"github.com/urfave/cli/v2"
)

func (runner *Runner) setCLIArg(c *cli.Context, param *controller.Param) error { //nolint:unparam
	param.StatePath = c.String("state")
	if logLevel := c.String("log-level"); logLevel != "" {
		param.LogLevel = logLevel
	}
	param.DryRun = c.Bool("dry-run")
	param.ConfigFilePath = c.String("config")
	if param.ConfigFilePath == "" {
		param.ConfigFilePath = "tfmigrator.yaml"
	}
	param.HCLFilePaths = c.Args().Slice()
	return nil
}

func (runner *Runner) runAction(c *cli.Context) error {
	param := &controller.Param{}
	if err := runner.setCLIArg(c, param); err != nil {
		return fmt.Errorf("parse the command line arguments: %w", err)
	}

	ctrl, param, err := controller.New(c.Context, param)
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}

	return ctrl.Run(c.Context, param) //nolint:wrapcheck
}
