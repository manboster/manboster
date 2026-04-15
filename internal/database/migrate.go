package database

import (
	"github.com/manboster/manboster/internal/database/types"
)

func (c *Client) Migrate() error {
	err := c.db.AutoMigrate(&types.User{})
	if err != nil {
		return err
	}

	err = c.db.AutoMigrate(&types.Chat{})
	if err != nil {
		return err
	}

	err = c.db.AutoMigrate(&types.Session{})
	if err != nil {
		return err
	}

	err = c.db.AutoMigrate(&types.Memory{})
	if err != nil {
		return err
	}

	err = c.db.AutoMigrate(&types.Soul{})
	if err != nil {
		return err
	}

	err = c.db.AutoMigrate(&types.ChatData{})
	if err != nil {
		return err
	}

	return nil
}
