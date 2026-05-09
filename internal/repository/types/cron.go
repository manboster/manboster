package types

import "github.com/manboster/manboster/internal/database/types"

type Cron struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	ChatID       string `json:"chat_id"`
	ChatProvider string `json:"chat_provider"`
	CronTab      string `json:"cron_tab"`
	Type         string `json:"type"`
	Ignore       string `json:"ignore"`
	Prompt       string `json:"prompt"`
	CreateBy     string `json:"create_by"`
}

func MapCron(c types.Cron) Cron {
	return Cron{
		ID:           c.ID,
		Name:         c.Name,
		ChatID:       c.ChatID,
		ChatProvider: c.ChatProvider,
		Type:         c.Type,
		Ignore:       c.Ignore,
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
		Ignore:       c.Ignore,
		CronTab:      c.CronTab,
		Prompt:       c.Prompt,
		CreatedBy:    c.CreateBy,
	}
}
