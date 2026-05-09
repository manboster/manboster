package cron

import "context"

func (s *Service) Start(ctx context.Context) error {
	s.cron.Start() // 启动

	<-ctx.Done()
	s.cron.Stop()
	return ctx.Err()
}

func (s *Service) Stop() error {
	s.cron.Stop()
	return nil
}
