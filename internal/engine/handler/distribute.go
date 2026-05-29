package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (h *Handler) DistributeFeedbackMsg(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload, err error, hachimiStatus hachimi.ResponseStatusType) error {
	var txt strings.Builder
	respMsg := msg.Clone()
	respMsg.Reply = nil
	respMsg.MessageType = chat.MessageText
	respMsg.Text = &chat.TextPayload{}

	switch hachimiStatus {
	case hachimi.ResponseStatusInspect:
		txt.WriteString(fmt.Sprintf("🐱❓"))
	case hachimi.ResponseStatusSafe:
		txt.WriteString(fmt.Sprintf("🐱✅"))
	case hachimi.ResponseStatusUnsafe:
		txt.WriteString(fmt.Sprintf("🐱❌"))
	default:
		txt.WriteString(fmt.Sprintf("🐱➖"))
	}

	if err == nil {
		txt.WriteString(" 🤖✅ ")
	} else {
		txt.WriteString(" 🤖❌ ")
	}

	if toolProvider.MetaData().Represent == "" {
		txt.WriteString(fmt.Sprintf("🧰 `%s`", toolProvider.DisplayName()))
	} else {
		txt.WriteString(fmt.Sprintf("%s `%s`", toolProvider.MetaData().Represent, toolProvider.DisplayName()))
	}

	if err == nil {
		params := toolProvider.ClientRenderer(fmt.Sprintf("%s", req.ToolArgs))
		if params != "" {
			txt.WriteString(fmt.Sprintf(": `%s`", params))
		}
	} else {
		txt.WriteString(fmt.Sprintf(": `%q`", err))
	}
	txt.WriteString("\n")

	count := h.sessionManager.Chat.GetToolCallCounts(sid)
	if count%10 == 0 {
		respMsg.Text.Text = txt.String()
		err = h.gateway.SendMessage(ctx, instance, respMsg)
		h.sessionManager.Chat.ResetTool(sid, respMsg.MessageID)
		h.sessionManager.Chat.SetToolMsgData(sid, txt.String())
		return err
	}

	msgId := h.sessionManager.Chat.GetToolMsgId(sid)
	data := h.sessionManager.Chat.GetToolMsgData(sid)
	respMsg.Text.Text = data + "\n" + txt.String()
	h.sessionManager.Chat.SetToolMsgData(sid, respMsg.Text.Text)

	respMsg.MessageType |= chat.MessageUnknown
	respMsg.MessageID = msgId
	return h.gateway.EditMessage(ctx, instance, respMsg)
}
