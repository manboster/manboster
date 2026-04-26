package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func (c *Client) Instance() *gorm.DB {
	return c.db
}

func (c *Client) Init(path string) error {
	dbi, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent), // shut up, sir
	})
	c.db = dbi
	if err != nil {
		return err
	}

	err = c.Migrate()
	if err != nil {
		return err
	}

	return nil
}

var DBInstance *Client
