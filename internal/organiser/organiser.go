package organiser

import (
	"file-organiser/internal/config"
	"file-organiser/internal/utils"
	"file-organiser/pkg/logger"
	"fmt"
	"os"
	"path/filepath"
)

type Organizer struct {
	cfg    *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, logger logger.Logger) *Organizer {
	return &Organizer{cfg: cfg, logger: logger}
}

func (o *Organizer) Organize() error {
	files, err := utils.ListFiles(o.cfg.SourceDir)
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	for _, file := range files {
		if err := o.organizeFile(file); err != nil {
			o.logger.Error("Failed to organize file", "file", file, "error", err)
		}
	}

	return nil
}

func (o *Organizer) organizeFile(file string) error {
	ext := filepath.Ext(file)
	targetDir := o.cfg.FileExtensions[ext]
	if targetDir == "" {
		targetDir = "misc"
	}

	targetPath := filepath.Join(o.cfg.SourceDir, targetDir)
	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	newPath := filepath.Join(targetPath, filepath.Base(file))
	if err := os.Rename(file, newPath); err != nil {
		return fmt.Errorf("failed to move file: %w", err)
	}

	o.logger.Info("File organized", "from", file, "to", newPath)
	return nil
}