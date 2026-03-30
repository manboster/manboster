package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Manboster: Your personal Manbo Lobster!
// Powered by chihuo2104(c)2026.
// Last Update: 2026.3.31

func main() {
	var rootCmd = &cobra.Command{
		Use:   "manboster",
		Short: "manboster：你的曼波虾头小助手",
		Long:  `manboster是一个由ai大语言模型驱动的ai助手，致力在守护你设备安全的情况下给你最佳的体验！`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World!")
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "查看manboster的版本",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("0.0.0")
		},
	}

	rootCmd.AddCommand(versionCmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
