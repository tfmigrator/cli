package config

import (
	"fmt"
	"os"

	"github.com/suzuki-shunsuke/expr-unmarshaler/expr"
	"github.com/suzuki-shunsuke/go-template-unmarshaler/text"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Rules []*Rule
}

type Rule struct {
	If                 *expr.Bool
	Address            *text.Template
	Dirname            *text.Template
	HCLFileBasename    *text.Template `yaml:"hcl_file_basename"`
	StateBasename      *text.Template `yaml:"state_file_basename"`
	Removed            bool
	SkipHCLMigration   bool `yaml:"skip_hcl_migration"`
	SkipStateMigration bool `yaml:"skip_state_migration"`
}

func Read(p string, cfg *Config) error {
	cfgFile, err := os.Open(p)
	if err != nil {
		return fmt.Errorf("open a configuration file %s: %w", p, err)
	}
	defer cfgFile.Close()
	if err := yaml.NewDecoder(cfgFile).Decode(&cfg); err != nil {
		return fmt.Errorf("parse a configuration file as YAML %s: %w", p, err)
	}
	return nil
}
