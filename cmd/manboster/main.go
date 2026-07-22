package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/manboster"
	"github.com/manboster/manboster/internal/i18n"
)

// Manboster: Your Personal AI Agent Manbo Lobster!
// Powered by chihuo2104 and the other Manboster contributors (c) 2026.
// Last Update: 2026.7.22

func main() {
	err := i18n.Init()
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster] Error initializing i18n: %v", err))
	}
	manboster.Init()
}
