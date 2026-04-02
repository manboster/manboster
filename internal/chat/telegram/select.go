package telegram

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

// Select TODO: give user a plenty of selections and wait for them to reply.
func (s *Service) Select(ctx context.Context, title string, name string, selection []chat.Selection) (string, error) {
	return "", nil
}
