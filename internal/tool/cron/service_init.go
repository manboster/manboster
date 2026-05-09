package cron

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/glebarez/sqlite"
	"github.com/manboster/manboster/internal/database"
	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (s *Service) Init(ctx context.Context, cfg any) error {
	if database.DBInstance != nil {
		dbi := database.DBInstance.Instance()
		err := dbi.AutoMigrate(dbtypes.Cron{})
		if err != nil {
			return err
		}
		s.cronRepo = repository.NewCronRepo(dbi)
	} else {
		// downgrade to memory storage
		color.Yellow(fmt.Sprintf("[Manboster Tool] dev.manboster.cron downgraded to memory repository, this session's storage is not persistent!"))
		dbi, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // shut up, sir
		})
		if err != nil {
			return err
		}
		err = dbi.AutoMigrate(&dbtypes.Cron{})
		if err != nil {
			return err
		}
		s.cronRepo = repository.NewCronRepo(dbi)
	}
	s.manager = NewManager()
	return nil
}
