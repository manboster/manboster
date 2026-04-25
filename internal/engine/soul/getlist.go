package soul

import (
	"context"
	"sort"
)

func (s *Service) GetSoulsList(ctx context.Context, chatId string) []string {
	seen := map[string]bool{"system": true}
	soulsList := []string{"system"}

	souls, err := s.repo.ReadSoulsByScope(ctx, "*")
	if err != nil {
		return soulsList
	}
	sort.Slice(souls, func(i, j int) bool {
		return souls[i].Priority < souls[j].Priority
	})
	for _, soul := range souls {
		if !seen[soul.Name] {
			soulsList = append(soulsList, soul.Name)
			seen[soul.Name] = true
		}
	}

	souls, err = s.repo.ReadSoulsByScope(ctx, "chat:"+chatId)
	if err != nil {
		return soulsList
	}
	sort.Slice(souls, func(i, j int) bool {
		return souls[i].Priority < souls[j].Priority
	})

	for _, soul := range souls {
		if !seen[soul.Name] {
			soulsList = append(soulsList, soul.Name)
			seen[soul.Name] = true
		}
	}
	return soulsList
}
