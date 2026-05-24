package memory_kv

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
	if memDB != nil {
		return nil
	}

	if database.DBInstance != nil {
		dbi := database.DBInstance.Instance()
		if err := dbi.AutoMigrate(dbtypes.Memory{}); err != nil {
			return err
		}
		memDB = repository.NewMemoryRepo(dbi)
	} else {
		color.Yellow(fmt.Sprintf("[Manboster Tool] dev.manboster.memory downgraded to memory repository, this session's storage is not persistent!"))
		dbi, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			return err
		}
		if err := dbi.AutoMigrate(&dbtypes.Memory{}); err != nil {
			return err
		}
		memDB = repository.NewMemoryRepo(dbi)
	}
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}
