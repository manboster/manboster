package skills

import (
	"context"

	"github.com/manboster/manboster/internal/engine/hook"
)

func (s *Service) RegisterHook(registry *hook.Registry) {
	registry.Register(hook.EngineBeforeBuildSystemPrompt, "skills-hook", hook.EngineBeforeBuildSystemPromptHookProvider{
		PolyfillFunc: s.InjectSkills,
	})
}

func (s *Service) InjectSkills(ctx context.Context, before string) (string, error) {
	return before, nil
}
