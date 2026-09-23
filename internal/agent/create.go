package agent

import "github.com/manboster/manboster/spec/llm"

// Create creates a father agent providing for a single session
func Create(provider llm.Provider, model llm.Model, session string) (*Agent, error) {
	return &Agent{
		Provider: provider,
		Powered:  model,
		Session:  session,
		Father:   nil,
	}, nil
}

// Create creates a child agent for an agent
func (a *Agent) Create(provider llm.Provider, model llm.Model, session string) (*Agent, error) {
	child := &Agent{
		Provider: provider,
		Powered:  model,
		Session:  session,
		Father:   a,
	}

	a.Children = append(a.Children, child)
	return child, nil
}
