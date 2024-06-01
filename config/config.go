package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type Workflow struct {
	SourceBranch      string `toml:"source_branch"`
	DestinationBranch string `toml:"destination_branch"`
	DefaultBranch     string `toml:"default_branch,omitempty"`
	PreCommitHooks    string `toml:"pre_commit_hooks,omitempty"`
	PostCommitHooks   string `toml:"post_commit_hooks,omitempty"`
}

type Config struct {
	Workflows map[string]Workflow
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{
		Workflows: make(map[string]Workflow),
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(file, &config.Workflows); err != nil {
		return nil, err
	}

	return config, nil
}
