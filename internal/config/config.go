package config

import (
	"flag"
	"fmt"
	"strings"

	"code.gitea.io/sdk/gitea"
)

type IssueState string

const (
	StateAll    IssueState = "all"
	StateOpen   IssueState = "open"
	StateClosed IssueState = "closed"
)

type Config struct {
	BaseURL   string
	Token     string
	Owner     string
	Repo      string
	OutputDir string
	State     IssueState
}

func Parse() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.BaseURL, "url", "https://gitea.com", "Gitea instance URL")
	flag.StringVar(&cfg.Token, "token", "", "Gitea access token")
	flag.StringVar(&cfg.OutputDir, "output", "issues", "output directory for markdown files")
	state := flag.String("state", "all", "issue state to export (all, open, closed)")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		return nil, fmt.Errorf("repository must be specified as owner/repo")
	}

	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid repository format, expected owner/repo")
	}

	cfg.Owner = parts[0]
	cfg.Repo = parts[1]

	switch *state {
	case "all", "open", "closed":
		cfg.State = IssueState(*state)
	default:
		return nil, fmt.Errorf("invalid state: %s (must be all, open, or closed)", *state)
	}

	if cfg.Token == "" {
		return nil, fmt.Errorf("token is required")
	}

	return cfg, nil
}

func (s IssueState) ToGiteaState() gitea.StateType {
	switch s {
	case StateOpen:
		return gitea.StateOpen
	case StateClosed:
		return gitea.StateClosed
	default:
		return gitea.StateAll
	}
}
