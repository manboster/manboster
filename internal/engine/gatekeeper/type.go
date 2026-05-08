package gatekeeper

import "github.com/manboster/manboster/spec/chat"

var selectionNoHachimi = []chat.Selection{
	{
		Name:  "Continue",
		Value: "continue",
	},
	{
		Name:  "Continue and shut up",
		Value: "hachimi",
	},
	{
		Name:  "Cancel",
		Value: "cancel",
	},
	{
		Name:  "Cancel and silence in 15 minutes",
		Value: "cAnCel",
	},
}

var selectionWithHachimi = []chat.Selection{
	{
		Name:  "Continue",
		Value: "continue",
	},
	{
		Name:  "Continue and shut up, handled by hachimi",
		Value: "hachimi",
	},
	{
		Name:  "Cancel",
		Value: "cancel",
	},
	{
		Name:  "Cancel and silence in 15 minutes",
		Value: "cAnCel",
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
