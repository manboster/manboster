package ignorance

import "time"

type mark struct {
	M          bool
	ActionTime time.Time
	Ttl        int
	MarkType   MarkType
}

type MarkType string

const (
	MarkCancel      MarkType = "cAnCel"
	MarkIgnore      MarkType = "ignore"
	MarkHachimi     MarkType = "hachimi"
	MarkHachimiAll  MarkType = "hachimi_all"
	MarkContinueAll MarkType = "ignore_all"
	MarkCancelAll   MarkType = "cancel_all"
)
