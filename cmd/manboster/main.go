package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/manboster"
	"github.com/manboster/manboster/internal/i18n"
)

// Manboster: Your Personal Manbo Lobster!
// Powered by the Manboster contributors (c) 2026.
// Last Update: 2026.6.20

func main() {
	err := i18n.Init()
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster] Error initializing i18n: %v", err))
	}
	manboster.Init()
}
