package memory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/glebarez/sqlite"
	"github.com/manboster/manboster/internal/database"
	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (s *Service) Init(ctx context.Context) error {
	if database.DBInstance != nil {
		dbi := database.DBInstance.Instance()
		err := dbi.AutoMigrate(dbtypes.Memory{})
		if err != nil {
			return err
		}
		s.memDB = repository.NewMemoryRepo(dbi)
	} else {
		// downgrade to memory storage
		color.Yellow(fmt.Sprintf("[Manboster Tool] dev.manboster.memory downgraded to memory repository, this session's storage is not persistent!"))
		dbi, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // shut up, sir
		})
		if err != nil {
			return err
		}
		err = dbi.AutoMigrate(&dbtypes.Memory{})
		if err != nil {
			return err
		}
		s.memDB = repository.NewMemoryRepo(dbi)
	}
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Run(ctx context.Context, args string) (string, error) {
	arg := RunArgs{}
	// fmt.Println(args)
	if json.Unmarshal([]byte(args), &arg) == nil {
		switch arg.Name {
		case "get":
			if arg.Key != "" {
				memory, err := s.memDB.GetMemory(ctx, arg.Key)
				if err != nil {
					return "", fmt.Errorf("failed to get %q", arg.Key)
				}
				return memory.Value, nil
			}
		case "set":
			if arg.Key != "" && arg.Value != "" {
				err := s.memDB.EditMemoryValue(ctx, arg.Key, arg.Value)
				if errors.Is(err, repository.ErrNotFound) {
					err = s.memDB.CreateMemory(ctx, types.Memory{
						Key:   arg.Key,
						Value: arg.Value,
					})
				}
				if err != nil {
					return "", fmt.Errorf("failed to storage %q", arg.Key)
				}
				return "success", nil
			}
		case "delete":
			if arg.Key != "" {
				err := s.memDB.DeleteMemory(ctx, arg.Key)
				if err != nil {
					return "", fmt.Errorf("failed to delete %q", arg.Key)
				}
				return "success", nil
			}
		case "list":
			keys, err := s.memDB.ListMemoryKeys(ctx)
			if err != nil {
				return "", fmt.Errorf("failed to list keys")
			}
			jsonify, err := json.Marshal(keys)
			if err != nil {
				return "", fmt.Errorf("failed to marshal keys")
			}
			return string(jsonify), nil
		default:
			return "", fmt.Errorf("unknown argument %q", arg.Name)
		}
	}
	return "", fmt.Errorf("invalid arguments")
}

func (s *Service) Close() error {
	return nil
}
