package soul

import (
	"context"
)

func (s *Service) Init(ctx context.Context) error {
	if s.repo == nil {
		return ErrNoRepositoryAvailable
	}
	data, err := s.repo.GetAllSouls(ctx)
	if err != nil {
		return err
	}
	for _, soul := range data {
		s.soulMap[soul.Name] = soul
	}
	return nil
}
