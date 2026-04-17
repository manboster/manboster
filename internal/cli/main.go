package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/loader"
	"github.com/spf13/cobra"

	_ "github.com/manboster/manboster/internal/chat/telegram"
	_ "github.com/manboster/manboster/internal/llm/oai_compat"
	_ "github.com/manboster/manboster/internal/llm/openrouter"
)

// main is the entrypoint function that when user runs 'manboster'.
func main(cmd *cobra.Command, args []string) {
	// output welcome
	color.Cyan("Welcome to Manboster!")
	color.Blue("[Manboster Client] Your Lobster is on the way, please wait...")

	// create a universal context for this application
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	color.Blue(fmt.Sprintf("[Manboster Client] Reading Configuration..."))
	cfg := config.Read()

	// create a loader instance
	loaderInstance := loader.New(&cfg)
	err := loaderInstance.Load(ctx)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Client] Error while load using the loader: %q", err))
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}

	<-ctx.Done()
	color.Yellow("[Manboster Client] Your Manboster is going to sleep, thank you for playing with it!")
	time.Sleep(1 * time.Second)
}
