package types

import "github.com/manboster/manboster/internal/database/types"

type Cron struct {
	ID           uint64
	Name         string
	ChatID       string
	ChatProvider string
	CronTab      string
	Type         string
	Prompt       string
	CreateBy     string
}

func MapCron(c types.Cron) Cron {
	return Cron{
		ID:           c.ID,
		Name:         c.Name,
		ChatID:       c.ChatID,
		ChatProvider: c.ChatProvider,
		Type:         c.Type,
		CronTab:      c.CronTab,
		Prompt:       c.Prompt,
		CreateBy:     c.CreatedBy,
	}
}

func MapCr(c Cron) types.Cron {
	return types.Cron{
		ID:           c.ID,
		Name:         c.Name,
		ChatID:       c.ChatID,
		ChatProvider: c.ChatProvider,
		Type:         c.Type,
		CronTab:      c.CronTab,
		Prompt:       c.Prompt,
		CreatedBy:    c.CreateBy,
	}
}
