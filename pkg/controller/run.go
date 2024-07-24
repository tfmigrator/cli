package controller

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/suzuki-shunsuke/go-template-unmarshaler/text"
	"github.com/tfmigrator/cli/pkg/config"
	"github.com/tfmigrator/cli/pkg/planner"
	"github.com/tfmigrator/tfmigrator/tfmigrator"
	"github.com/tfmigrator/tfmigrator/tfmigrator/hcledit"
	tflog "github.com/tfmigrator/tfmigrator/tfmigrator/log"
	"github.com/tfmigrator/tfmigrator/tfmigrator/tfstate"
)

func (ctrl *Controller) Run(ctx context.Context, param *Param) error { //nolint:funlen
	text.SetTemplateFunc(func(s string) (*template.Template, error) {
		return template.New("_").Funcs(sprig.TxtFuncMap()).Parse(s)
	})
	cfg := config.Config{}
	if err := config.Read(param.ConfigFilePath, &cfg); err != nil {
		return fmt.Errorf("read the configuration file %s: %w", param.ConfigFilePath, err)
	}
	pln := &planner.Planner{
		Rules: cfg.Rules,
	}

	logger := &tflog.SimpleLogger{}
	if param.LogLevel != "" {
		if err := logger.SetLogLevel(param.LogLevel); err != nil {
			return fmt.Errorf("set the log level (%s): %w", param.LogLevel, err)
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get the current directory: %w", err)
	}
	tfCmdPath, err := exec.LookPath("terraform")
	if err != nil {
		return errors.New("the command `terraform` isn't found: %w")
	}
	tf, err := tfexec.NewTerraform(wd, tfCmdPath)
	if err != nil {
		return fmt.Errorf("initialize Terraform exec: %w", err)
	}

	editor := &hcledit.Client{
		DryRun: param.DryRun,
		Stderr: os.Stderr,
		Logger: logger,
	}

	runner := &tfmigrator.Runner{
		Planner: pln,
		Logger:  logger,
		HCLEdit: editor,
		StateReader: &tfstate.Reader{
			Stderr:    os.Stderr,
			Logger:    logger,
			Terraform: tf,
		},
		Outputter: tfmigrator.NewYAMLOutputter(os.Stderr),
		Migrator: &tfmigrator.Migrator{
			Stdout:  os.Stdout,
			DryRun:  param.DryRun,
			HCLEdit: editor,
			StateUpdater: &tfstate.Updater{
				Stdout:    os.Stdout,
				Stderr:    os.Stderr,
				DryRun:    param.DryRun,
				Logger:    logger,
				Terraform: tf,
			},
		},
		DryRun: param.DryRun,
	}
	return runner.Run(ctx, &tfmigrator.RunOpt{ //nolint:wrapcheck
		SourceHCLFilePaths: param.HCLFilePaths,
		SourceStatePath:    param.StatePath,
	})
}
