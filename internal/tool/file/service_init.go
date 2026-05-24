package file

import (
	"context"
	"os"
	"path/filepath"

	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/config"
)

func (s *Service) Init(ctx context.Context, conf any) error {
	if cfg != nil {
		return nil
	}

	err := os.MkdirAll(config.Path("workspace"), 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(config.Path(filepath.Join("workspace", "public")), 0755)
	if err != nil {
		return err
	}

	var c Config
	err = mapstructure.Decode(conf, &c)
	if err != nil {
		return err
	}

	cfg = &c
	return cfg.Validate()
}
