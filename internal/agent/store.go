package agent

type Store struct {
	ID       string   `json:"id"`
	Data     Agent    `json:"data"`
	Father   string   `json:"father"`
	Children []string `json:"children"`
}

// ReStore restores Agent type from the database Store
func (s Store) ReStore() Agent {
	return s.Data // TODO: read stored things in database and then convert them into agent
}
