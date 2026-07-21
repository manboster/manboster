package agent

import "github.com/manboster/manboster/spec/llm"

// Agent marks a new era of Manboster agenting
type Agent struct {
	ID       string       // Agent ID
	Powered  llm.Model    // Powered by which model
	Provider llm.Provider // Provided by who
	Events   []llm.Event  // Events used
	Cost     llm.Usage    // Usage used by this agent
	Session  string       // SessionID belongs to
	Father   *Agent       // If it's nil, it's the root agent of Manboster
	Children []*Agent     // The Agent's children
}
