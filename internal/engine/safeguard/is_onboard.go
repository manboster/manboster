package safeguard

func (s *Service) IsOnboarding() bool {
	return s.onboard != nil && s.onboard.Active()
}
