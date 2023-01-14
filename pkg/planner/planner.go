package planner

import (
	"fmt"

	"github.com/tfmigrator/cli/pkg/config"
	"github.com/tfmigrator/tfmigrator/tfmigrator"
)

type Planner struct {
	Rules []*config.Rule
}

func (planner *Planner) Plan(src *tfmigrator.Source) (*tfmigrator.MigratedResource, error) {
	for _, rule := range planner.Rules {
		matched, err := rule.If.Run(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the rule: %w", err)
		}
		if !matched {
			continue
		}
		if rule.Ignored {
			return nil, nil //nolint:nilnil
		}
		return planner.plan(src, rule)
	}
	return nil, nil //nolint:nilnil
}

func (planner *Planner) plan(src *tfmigrator.Source, rule *config.Rule) (*tfmigrator.MigratedResource, error) {
	rsc := &tfmigrator.MigratedResource{
		Removed:            rule.Removed,
		SkipHCLMigration:   rule.SkipHCLMigration,
		SkipStateMigration: rule.SkipStateMigration,
	}
	if rule.Removed {
		return rsc, nil
	}

	if !rule.Address.Empty() {
		s, err := rule.Address.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the address: %w", err)
		}
		rsc.Address = s
	}

	if !rule.Dirname.Empty() {
		s, err := rule.Dirname.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the dirname: %w", err)
		}
		rsc.Dirname = s
	}

	if !rule.HCLFileBasename.Empty() {
		s, err := rule.HCLFileBasename.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the hcl_file_basename: %w", err)
		}
		rsc.HCLFileBasename = s
	}

	if !rule.StateBasename.Empty() {
		s, err := rule.StateBasename.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the state_basename: %w", err)
		}
		rsc.StateBasename = s
	}

	return rsc, nil
}
