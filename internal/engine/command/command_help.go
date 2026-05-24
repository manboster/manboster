package command

import (
	"context"
	"strings"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// cmdHelp is the helper command of manboster
func (h *Handler) cmdHelp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	var str strings.Builder
	str.WriteString(i18n.T(keys.CmdHelpHeader))
	str.WriteString(i18n.T(keys.CmdHelpVersion))
	str.WriteString(i18n.T(keys.CmdHelpID))
	str.WriteString(i18n.T(keys.CmdHelpHelp))
	str.WriteString(i18n.T(keys.CmdHelpOp))
	str.WriteString(i18n.T(keys.CmdHelpDeop))
	str.WriteString(i18n.T(keys.CmdHelpStatus))
	str.WriteString(i18n.T(keys.CmdHelpSave))
	str.WriteString(i18n.T(keys.CmdHelpNew))
	str.WriteString(i18n.T(keys.CmdHelpCompact))
	str.WriteString(i18n.T(keys.CmdHelpModel))
	str.WriteString(i18n.T(keys.CmdHelpModels))
	str.WriteString(i18n.T(keys.CmdHelpSession))
	str.WriteString(i18n.T(keys.CmdHelpSessions))
	str.WriteString(i18n.T(keys.CmdHelpProvider))
	str.WriteString(i18n.T(keys.CmdHelpProviders))
	str.WriteString(i18n.T(keys.CmdHelpStart))
	str.WriteString(i18n.T(keys.CmdHelpPair))
	str.WriteString(i18n.T(keys.CmdHelpCancel))

	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: str.String(),
	}
	return instance.SendMessage(ctx, msg)
}
