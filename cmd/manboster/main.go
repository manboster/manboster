package main

import (
	"github.com/manboster/manboster/internal/cli"
	"github.com/manboster/manboster/internal/config"
)

// Manboster: Your Personal Manbo Lobster!
// Powered by chihuo2104 (c) 2026.
// Last Update: 2026.4.1

func main() {
	config.Init()
	cli.Init()
}
