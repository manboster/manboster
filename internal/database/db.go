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
	dbi, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
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
