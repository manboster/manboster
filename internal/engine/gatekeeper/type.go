package gatekeeper

import (
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

type guardSelectType string

const (
	guardSelectCancel       guardSelectType = "cancel"
	guardSelectCancelIgnore guardSelectType = "cAnCel"
	guardSelectContinue     guardSelectType = "continue"
	guardSelectHachimi      guardSelectType = "hachimi"
	guardSelectHachimiAll   guardSelectType = "hachimi_all"
	guardSelectIgnore       guardSelectType = "ignore"
	guardSelectContinueAll  guardSelectType = "continue_all"
	guardSelectCancelAll    guardSelectType = "cancel_all"
)

var selectionNoHachimi = []chat.Selection{
	{
		Name:  i18n.T(keys.GatekeeperContinueOnce),
		Value: string(guardSelectContinue),
	},
	{
		Name:  i18n.T(keys.GatekeeperContinueAll),
		Value: string(guardSelectContinueAll),
	},
	{
		Name:  i18n.T(keys.GatekeeperShutUp),
		Value: string(guardSelectIgnore),
	},
	{
		Name:  i18n.T(keys.GatekeeperCancelOnce),
		Value: string(guardSelectCancel),
	},
	{
		Name:  i18n.T(keys.GatekeeperCancelIgnore),
		Value: string(guardSelectCancelIgnore),
	},
	{
		Name:  i18n.T(keys.GatekeeperCancelAll),
		Value: string(guardSelectCancelAll),
	},
}

var selectionWithHachimi = []chat.Selection{
	{
		Name:  i18n.T(keys.GatekeeperContinueOnce),
		Value: string(guardSelectContinue),
	},
	{
		Name:  i18n.T(keys.GatekeeperContinueAll),
		Value: string(guardSelectContinueAll),
	},
	{
		Name:  i18n.T(keys.GatekeeperHandleHachimi),
		Value: string(guardSelectHachimi),
	},
	{
		Name:  i18n.T(keys.GatekeeperHandleHachimiAll),
		Value: string(guardSelectHachimiAll),
	},
	{
		Name:  i18n.T(keys.GatekeeperShutUp),
		Value: string(guardSelectIgnore),
	},
	{
		Name:  i18n.T(keys.GatekeeperCancelOnce),
		Value: string(guardSelectCancel),
	},
	{
		Name:  i18n.T(keys.GatekeeperCancelIgnore),
		Value: string(guardSelectCancelIgnore),
	},
	{
		Name:  i18n.T(keys.GatekeeperCancelAll),
		Value: string(guardSelectCancelAll),
	},
}

var selectionHachimi = []chat.Selection{
	{
		Name:  i18n.T(keys.GatekeeperAllow),
		Value: "allow",
	},
	{
		Name:  i18n.T(keys.GatekeeperDeny),
		Value: "deny",
	},
}
