package planner

import (
	"fmt"

	"github.com/tfmigrator/cli/pkg/config"
	"github.com/tfmigrator/tfmigrator/tfmigrator"
)

type Planner struct {
	Items []*config.Item
}

func (planner *Planner) Plan(src *tfmigrator.Source) (*tfmigrator.MigratedResource, error) {
	for _, item := range planner.Items {
		matched, err := item.Rule.Run(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the rule: %w", err)
		}
		if !matched {
			continue
		}
		return planner.plan(src, item)
	}
	return nil, nil
}

func (planner *Planner) plan(src *tfmigrator.Source, item *config.Item) (*tfmigrator.MigratedResource, error) {
	rsc := &tfmigrator.MigratedResource{
		Removed:            item.Removed,
		SkipHCLMigration:   item.SkipHCLMigration,
		SkipStateMigration: item.SkipStateMigration,
	}
	if item.Removed {
		return rsc, nil
	}

	if !item.Address.Empty() {
		s, err := item.Address.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the address: %w", err)
		}
		rsc.Address = s
	}

	if !item.Dirname.Empty() {
		s, err := item.Dirname.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the dirname: %w", err)
		}
		rsc.Dirname = s
	}

	if !item.HCLFileBasename.Empty() {
		s, err := item.HCLFileBasename.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the hcl_file_basename: %w", err)
		}
		rsc.HCLFileBasename = s
	}

	if !item.StateBasename.Empty() {
		s, err := item.StateBasename.Execute(src)
		if err != nil {
			return nil, fmt.Errorf("evaluate the state_basename: %w", err)
		}
		rsc.StateBasename = s
	}

	return rsc, nil
}
