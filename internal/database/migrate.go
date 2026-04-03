package database

import (
	"github.com/manboster/manboster/internal/database/types"
)

func (c *Client) Migrate() error {
	err := c.db.AutoMigrate(&types.User{})
	if err != nil {
		return err
	}
	
	return nil
}
