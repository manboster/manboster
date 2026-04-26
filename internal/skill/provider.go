package skill

import (
	"github.com/manboster/manboster/spec/plugin"
)

type Provider interface {
	plugin.Provider
	Register() RegisterData
}

type RegisterData struct {
	Command  bool   // can this skill append to the command?
	Cron     bool   // can this skill use cron job?
	Schedule string // schedule expression e.g. "0 * * * *"
}
