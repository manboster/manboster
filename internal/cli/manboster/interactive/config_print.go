package interactive

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/tool"
)

func (s *databaseConfigService) printConfigDatabaseSessionList(ctx context.Context) error {
	// string is sessionID
	chatsMap := make(map[string][]types.Chat)

	sessions, err := s.repo.GetSessions(ctx)
	if err != nil {
		return err
	}
	s.sessions = sessions

	chats, err := s.repo.GetAllChats(ctx)
	if err != nil {
		return err
	}
	s.chats = chats

	for _, c := range chats {
		cm, avail := chatsMap[c.SessionID]
		if !avail {
			cm = []types.Chat{c}
			chatsMap[c.SessionID] = cm
			continue
		}
		cm = append(cm, c)
		chatsMap[c.SessionID] = cm
	}
	s.chatsMap = chatsMap

	var outputMsg strings.Builder
	for _, sess := range s.sessions {
		outputMsg.WriteString(fmt.Sprintf("%d) `%s`, used `%s:%s` created at %s, updated at %s.\n", sess.ID, sess.SessionID, sess.LLMProvider, sess.LLMProviderModel, sess.CreatedAt.Format("2006-01-02 15:04:05"), sess.UpdatedAt.Format("2006-01-02 15:04:05")))
		cm, avail := chatsMap[sess.SessionID]
		if avail {
			outputMsg.WriteString(fmt.Sprintf("Bind %d chats: ", len(cm)))
			for _, c := range cm {
				outputMsg.WriteString(fmt.Sprintf("%s:%s ", c.ChatProvider, c.ChatID))
			}
			outputMsg.WriteString("\n")
		}
	}

	outputMsg.WriteString(fmt.Sprintf("%d sessions loaded, %d sessions can be purged.", len(s.sessions), len(sessions)-len(chatsMap)))

	helper.DisplayText(outputMsg.String())
	return nil
}

func printConfigChatProvidersData(ctx context.Context) {
	var outputMsg strings.Builder
	conf := config.Read()
	for i, cp := range conf.Chats {
		inst, err := chat.GetProvider(cp.Provider)
		if err != nil {
			outputMsg.WriteString(fmt.Sprintf("%d) Could not get this!\n", i+1))
			continue
		}
		cfg := inst.Config()
		// get config
		err = mapstructure.Decode(cp.Configuration, &cfg)
		if err != nil {
			outputMsg.WriteString(fmt.Sprintf("%d) Could not get this!\n", i+1))
			continue
		}
		outputMsg.WriteString(fmt.Sprintf("%d) `%s`, config: %s\n", i+1, inst.DisplayName(), cfg))
	}
	outputMsg.WriteString(fmt.Sprintf("`%d` Chat providers loaded.\n", len(conf.Chats)))
	helper.DisplayText(outputMsg.String())
}

func printConfigLLMProvidersData(ctx context.Context) {
	var outputMsg strings.Builder
	conf := config.Read()
	for i, cp := range conf.LLMs {
		inst, err := llm.GetProvider(cp.Provider)
		if err != nil {
			outputMsg.WriteString(fmt.Sprintf("%d) Could not get this!\n", i+1))
			continue
		}
		cfg := inst.Config()
		// get config
		err = mapstructure.Decode(cp.Configuration, &cfg)
		if err != nil {
			outputMsg.WriteString(fmt.Sprintf("%d) Could not get this!\n", i+1))
			continue
		}
		err = inst.Init(ctx, cp.Configuration)
		if err != nil {
			outputMsg.WriteString(fmt.Sprintf("%d) Could not get this!\n", i+1))
			continue
		}
		outputMsg.WriteString(fmt.Sprintf("%d) `%s`, config: %s\n", i+1, inst.DisplayName(), cfg))
	}
	outputMsg.WriteString(fmt.Sprintf("%d LLM providers loaded.\n", len(conf.LLMs)))
	helper.DisplayText(outputMsg.String())
}

func printConfigHachimiProvidersData(ctx context.Context) {
	// TODO: wait for hachimi ends
}

func printConfigToolProvidersData(ctx context.Context) {
	var outputMsg strings.Builder
	conf := config.Read()
	for i, cp := range conf.Tools {
		inst, err := tool.GetProvider(cp.Name)
		if err != nil {
			outputMsg.WriteString(fmt.Sprintf("%d) Could not get this!\n", i+1))
			continue
		}

		outputMsg.WriteString(fmt.Sprintf("%d) `%s`", i+1, inst.DisplayName()))
		cfg := inst.Config()
		if cp.Configuration != nil {
			// get config
			err = mapstructure.Decode(cp.Configuration, &cfg)
			if err != nil {
				outputMsg.WriteString(fmt.Sprintf(" could not get this!\n"))
				continue
			}
			outputMsg.WriteString(fmt.Sprintf(", config: %s", cfg))
		}
		outputMsg.WriteString(fmt.Sprintf("\n"))
	}
	outputMsg.WriteString(fmt.Sprintf("%d Tool providers loaded.\n", len(conf.Tools)))
	helper.DisplayText(outputMsg.String())
}
