package cli

import (
	"context"
	"io"

	"github.com/tfmigrator/cli/pkg/constant"
	"github.com/urfave/cli/v2"
)

type Runner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (runner *Runner) Run(ctx context.Context, args ...string) error {
	app := cli.App{
		Name:    "tfmigrator",
		Usage:   "Migrate Terraform Configuration and State. https://github.com/tfmigrator/cli",
		Version: constant.Version,
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Migrate Terraform Configuration and State",
				Action: runner.runAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "log-level",
						Usage: "log level",
					},
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "configuration file path",
					},
					&cli.BoolFlag{
						Name:  "dry-run",
						Usage: "dry run",
					},
					&cli.StringFlag{
						Name:  "state",
						Usage: "the output of 'terraform show -json'",
					},
				},
			},
		},
	}

	return app.RunContext(ctx, args) //nolint:wrapcheck
}
