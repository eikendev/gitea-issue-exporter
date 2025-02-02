package main

import (
	"log/slog"
	"os"

	"github.com/eikendev/gitea-issue-exporter/internal/config"
	"github.com/eikendev/gitea-issue-exporter/internal/export"

	"code.gitea.io/sdk/gitea"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	cfg, err := config.Parse()
	if err != nil {
		logger.Error("failed to parse config", "error", err)
		os.Exit(1)
	}

	client, err := gitea.NewClient(cfg.BaseURL, gitea.SetToken(cfg.Token))
	if err != nil {
		logger.Error("failed to create client", "error", err)
		os.Exit(1)
	}

	exporter, err := export.New(client, cfg.OutputDir)
	if err != nil {
		logger.Error("failed to create exporter", "error", err)
		os.Exit(1)
	}

	if err := exporter.Export(cfg.Owner, cfg.Repo, cfg.State.ToGiteaState()); err != nil {
		logger.Error("failed to export issues", "error", err)
		os.Exit(1)
	}

	logger.Info("export complete", "output", cfg.OutputDir)
}
