package command

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

func (h *Handler) Init() {
	h.provider.Register(chat.CommandVersion, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdVersion(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandId, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdId(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandHelp, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdHelp(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandOp, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdOp(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandDeOp, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdDeOp(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandStatus, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdStatus)
	})

	h.provider.Register(chat.CommandSave, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		err := h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdCancel)
		if err != nil {
			return err
		}
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdSave)
	})

	h.provider.Register(chat.CommandNew, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		err := h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdCancel)
		if err != nil {
			return err
		}
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdNew)
	})

	h.provider.Register(chat.CommandCompact, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		err := h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdCancel)
		if err != nil {
			return err
		}
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.handler.HandleCompact)
	})

	h.provider.Register(chat.CommandModel, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdModel)
	})

	h.provider.Register(chat.CommandModels, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return nil // TODO: WIP
	})

	h.provider.Register(chat.CommandProvider, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdProvider)
	})

	h.provider.Register(chat.CommandProviders, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return nil // TODO: WIP
	})

	h.provider.Register(chat.CommandSession, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.handleAdminCommand(ctx, instance, msg, h.cmdSession)
	})

	h.provider.Register(chat.CommandSessions, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return nil // TODO: WIP
	})

	h.provider.Register(chat.CommandStart, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdStart(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandPair, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdPair(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandCancel, func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		return h.cmdCancel(ctx, instance, msg, sessionId)
	})

	h.provider.Default(func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
		if msg.ChatType == chat.ChatsPersonal {
			return h.cmdDefault(ctx, instance, msg)
		}
		return nil
	})
}
