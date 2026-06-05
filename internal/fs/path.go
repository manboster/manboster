package manbofs

import (
	"os"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
)

var path = config.Path("cache")

func init() {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		color.Yellow("[ManboFS] Could not create cache directory")
	}
}
