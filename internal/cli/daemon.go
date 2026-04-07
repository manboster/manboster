package cli

import "github.com/spf13/cobra"

func startCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the Manboster daemon in background",
		Run:   startCommandExecutor,
	}
}

func stopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop the Manboster daemon in background",
		Run:   stopCommandExecutor,
	}
}

func restartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "restart",
		Short: "Restart the Manboster daemon in background",
		Run:   restartCommandExecutor,
	}
}

func statusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the status of Manboster daemon in background",
		Run:   statusCommandExecutor,
	}
}
