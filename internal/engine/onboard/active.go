package onboard

func (s *Service) Active() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.active
}

func (s *Service) Deactivate() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.active = false
}
