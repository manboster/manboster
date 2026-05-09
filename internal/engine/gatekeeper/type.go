package gatekeeper

import "github.com/manboster/manboster/spec/chat"

type guardSelectType string

const (
	guardSelectCancel       guardSelectType = "cancel"
	guardSelectCancelIgnore guardSelectType = "cAnCel"
	guardSelectContinue     guardSelectType = "continue"
	guardSelectHachimi      guardSelectType = "hachimi"
	guardSelectIgnore       guardSelectType = "ignore"
)

var selectionNoHachimi = []chat.Selection{
	{
		Name:  "Continue",
		Value: string(guardSelectContinue),
	},
	{
		Name:  "Continue and shut up",
		Value: string(guardSelectIgnore),
	},
	{
		Name:  "Cancel",
		Value: string(guardSelectCancel),
	},
	{
		Name:  "Cancel and silence in 15 minutes",
		Value: string(guardSelectCancelIgnore),
	},
}

var selectionWithHachimi = []chat.Selection{
	{
		Name:  "Continue",
		Value: string(guardSelectContinue),
	},
	{
		Name:  "Continue and shut up, handled by hachimi",
		Value: string(guardSelectHachimi),
	},
	{
		Name:  "Continue and shut up",
		Value: string(guardSelectIgnore),
	},
	{
		Name:  "Cancel",
		Value: string(guardSelectCancel),
	},
	{
		Name:  "Cancel and silence in 15 minutes",
		Value: string(guardSelectCancelIgnore),
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
