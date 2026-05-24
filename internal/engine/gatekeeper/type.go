package gatekeeper

import "github.com/manboster/manboster/spec/chat"

type guardSelectType string

const (
	guardSelectCancel       guardSelectType = "cancel"
	guardSelectCancelIgnore guardSelectType = "cAnCel"
	guardSelectContinue     guardSelectType = "continue"
	guardSelectHachimi      guardSelectType = "hachimi"
	guardSelectHachimiAll   guardSelectType = "hachimi_all"
	guardSelectIgnore       guardSelectType = "ignore"
	guardSelectContinueAll  guardSelectType = "continue_all"
	guardSelectIgnoreAll    guardSelectType = "ignore_all"
)

var selectionNoHachimi = []chat.Selection{
	{
		Name:  "Continue this time",
		Value: string(guardSelectContinue),
	},
	{
		Name:  "Continue all in 10 minutes",
		Value: string(guardSelectContinueAll),
	},
	{
		Name:  "Shut up in a moment",
		Value: string(guardSelectIgnore),
	},
	{
		Name:  "Cancel this time",
		Value: string(guardSelectCancel),
	},
	{
		Name:  "Cancel in 15 minutes",
		Value: string(guardSelectCancelIgnore),
	},
	{
		Name:  "Cancel all in 10 minutes",
		Value: string(guardSelectIgnoreAll),
	},
}

var selectionWithHachimi = []chat.Selection{
	{
		Name:  "Continue this time",
		Value: string(guardSelectContinue),
	},
	{
		Name:  "Continue all in 10 minutes",
		Value: string(guardSelectContinueAll),
	},
	{
		Name:  "Handle to hachimi",
		Value: string(guardSelectHachimi),
	},
	{
		Name:  "Handle all to hachimi in an hour",
		Value: string(guardSelectHachimiAll),
	},
	{
		Name:  "Shut up in a moment",
		Value: string(guardSelectIgnore),
	},
	{
		Name:  "Cancel this time",
		Value: string(guardSelectCancel),
	},
	{
		Name:  "Cancel in 15 minutes",
		Value: string(guardSelectCancelIgnore),
	},
	{
		Name:  "Cancel all in 10 minutes",
		Value: string(guardSelectIgnoreAll),
	},
}

var selectionHachimi = []chat.Selection{
	{
		Name:  "Allow",
		Value: "allow",
	},
	{
		Name:  "Deny",
		Value: "deny",
	},
}
