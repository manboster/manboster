package command

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// Handle handles commands occurs
func (h *Handler) Handle(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	// TODO: check user
	if msg.Command == nil {
		return ErrInvalidParams
	}
	color.Blue(fmt.Sprintf("[Manboster Command Handler] Handling command... Received command: %s, args: %s", msg.Command.CommandType, msg.Command.CommandArgs))
	cType := msg.Command.CommandType

	h.provider.Register(chat.CommandVersion, func(ctx context.Context) error {
		return h.cmdVersion(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandId, func(ctx context.Context) error {
		return h.cmdId(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandHelp, func(ctx context.Context) error {
		return h.cmdHelp(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandOp, func(ctx context.Context) error {
		return h.cmdOp(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandDeOp, func(ctx context.Context) error {
		return h.cmdDeOp(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandStatus, func(ctx context.Context) error {
		return h.cmdStart(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandSave, func(ctx context.Context) error {
		err := h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdCancel)
		if err != nil {
			return err
		}
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdSave)
	})

	h.provider.Register(chat.CommandNew, func(ctx context.Context) error {
		err := h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdCancel)
		if err != nil {
			return err
		}
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdNew)
	})

	h.provider.Register(chat.CommandCompact, func(ctx context.Context) error {
		err := h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdCancel)
		if err != nil {
			return err
		}
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.handler.HandleCompact)
	})

	h.provider.Register(chat.CommandModel, func(ctx context.Context) error {
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdModel)
	})

	h.provider.Register(chat.CommandModels, func(ctx context.Context) error {
		return nil // TODO: WIP
	})

	h.provider.Register(chat.CommandProvider, func(ctx context.Context) error {
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdProvider)
	})

	h.provider.Register(chat.CommandProviders, func(ctx context.Context) error {
		return nil // TODO: WIP
	})

	h.provider.Register(chat.CommandSession, func(ctx context.Context) error {
		return h.handleAdminCommand(ctx, instance, msg, h.cmdSession)
	})

	h.provider.Register(chat.CommandSessions, func(ctx context.Context) error {
		return nil // TODO: WIP
	})

	h.provider.Register(chat.CommandStart, func(ctx context.Context) error {
		return h.cmdStart(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandPair, func(ctx context.Context) error {
		return h.cmdPair(ctx, instance, msg)
	})

	h.provider.Register(chat.CommandCancel, func(ctx context.Context) error {
		return h.cmdCancel(ctx, instance, msg, sessionId)
	})

	h.provider.Default(func(ctx context.Context) error {
		if msg.ChatType == chat.ChatsPersonal {
			return h.cmdDefault(ctx, instance, msg)
		}
		return nil
	})

	return h.provider.Handle(ctx, cType)
}

func (h *Handler) handleAdminCommand(ctx context.Context, instance chat.Provider, msg *chat.Message, call func(ctx context.Context, instance chat.Provider, msg *chat.Message) error) error {
	isAdmin := h.safeguardService.IsAdmin(h.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if isAdmin {
		return call(ctx, instance, msg)
	}
	return h.handler.HandleReject(ctx, instance, msg)
}

func (h *Handler) handleAdminCommandWithSessionID(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string, call func(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string) error) error {
	isAdmin := h.safeguardService.IsAdmin(h.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if isAdmin {
		return call(ctx, instance, msg, sid)
	}
	return h.handler.HandleReject(ctx, instance, msg)
}
