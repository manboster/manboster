package skill

import "context"

// Provider defines what a tool will be
type Provider interface {
	Name() string
	Description() string
	Requires() []string
	Register() RegisterData
	Args() map[string]any
	Run(ctx context.Context, args string) (string, error) // passthrough by JSON
}

type RegisterData struct {
	Command  bool   // can this skill append to the command?
	Cron     bool   // can this skill use cron job?
	Schedule string // schedule expression e.g. "0 * * * *"
}

type RequirementData struct {
	Type    RequirementType
	Package string
}

type RequirementType int16

const (
	RequireExecutable RequirementType = iota // executable file
	RequireTool                              // first find in tool, then in plugin
	RequireSkill                             // skills required to get
)
