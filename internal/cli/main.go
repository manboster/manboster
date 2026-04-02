package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/session"
	"github.com/spf13/cobra"

	_ "github.com/manboster/manboster/internal/chat/telegram"
	_ "github.com/manboster/manboster/internal/llm/oai_compat"
	_ "github.com/manboster/manboster/internal/llm/openrouter"
)

// main is the entrypoint function that when user runs 'manboster'.
func main(cmd *cobra.Command, args []string) {
	// initialize variables
	sessionManager := session.NewManager()

	// output welcome
	color.Cyan("Welcome to Manboster!")
	color.Blue("Your Lobster is on the way, please wait...")

	cfg := config.Read()
	err := cfg.Validate()
	if err != nil {
		os.Exit(1)
	}

	// create a universal context for this application
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// TODO: get model data from SQLite(Repository)
	// First, we activate LLMs.
	llmProviders := make([]llm.Provider, 0, len(cfg.LLMs))
	// configure and init LLM providers
	for _, llmConfigs := range cfg.LLMs {
		lProvider, err := llm.GetProvider(llmConfigs.Provider)
		if err != nil {
			color.Red(fmt.Sprintf("There is no provider named %q when importing chats providers. Please check your configuration.", llmConfigs.Provider))
			os.Exit(1)
		}
		newLProvider := lProvider.New()
		err = newLProvider.Init(ctx, llmConfigs.Configuration)
		if err != nil {
			color.Red("Activate ", lProvider.Name(), " Chat API Error! Message:", err.Error())
		}

		// append it into array!
		llmProviders = append(llmProviders, newLProvider)
	}

	// Then, we activate chats.
	for _, chatConfig := range cfg.Chats {
		cProvider, err := chat.GetProvider(chatConfig.Provider)
		if err != nil {
			os.Exit(1)
		}

		go func(instance chat.Provider, conf any) {
			err := instance.Start(ctx, conf, func(msg *chat.Message) {
				color.Blue(fmt.Sprintf("Got an message from %s by %s(%s)", instance.Name(), msg.Username, msg.UserID))
				sessionId := fmt.Sprintf("%s:%s", instance.Name(), msg.ChatID)
				sessionData := sessionManager.GetSession(sessionId)
				if len(sessionData.Messages) == 0 {
					sessionData.Messages = append(sessionData.Messages, llm.Message{
						Role: llm.RoleTypeSystem,
						Text: "You're an assistant named Manboster. You are chatting with people. The one who is chatting with you is your owner.", // TODO: prompt engineering
						Type: llm.MessageTypeText,
					})
				}
				msgData := append(sessionData.Messages, llm.Message{
					Role: llm.RoleTypeUser,
					Text: msg.Text,
					Type: llm.MessageTypeText,
				})

				tries := 0
				var mesg *llm.Message
				var err error
				// try 3 times
				for tries < 3 {
					mesg, err = llmProviders[0].Chat(ctx, msgData)
					if err != nil {
						color.Red("Retry ", tries, " times. Failed to get message from LLMProvider ", llmProviders[0].Name(), " Error:", err.Error())
						tries++
					} else {
						break
					}
				}
				if err != nil {
					color.Red("Failed to get message from LLMProvider ", llmProviders[0].Name(), " Error:", err.Error())
					msg.Text = "[Manboster]Failed to get message from LLMProvider " + llmProviders[0].Name() + ": " + err.Error()
				} else {
					msg.Text = mesg.Text
					msgData = append(msgData, llm.Message{
						Text: mesg.Text,
						Role: mesg.Role,
						Type: llm.MessageTypeText,
					})
				}

				sessionData.Messages = msgData
				sessionManager.SetSession(sessionId, sessionData)

				err = instance.SendMessage(ctx, msg)
				if err != nil {
					color.Red(err.Error())
					return
				}
			})
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
		}(cProvider, chatConfig.Configuration)
	}

	<-ctx.Done()
	color.Red("Your Manboster is going to sleep, thank you for playing with it!")
}
