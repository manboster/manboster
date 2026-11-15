package agent

import "github.com/manboster/manboster/spec/llm"

// Agent marks a new era of Manboster agenting
type Agent struct {
	ID       string       `json:"id"`       // Agent ID
	Powered  llm.Model    `json:"powered"`  // Powered by which model
	Provider llm.Provider `json:"provider"` // Provided by who
	Events   []string     `json:"events"`   // Events used by this agent
	Cost     llm.Usage    `json:"cost"`     // Usage used by this agent
	Session  string       `json:"session"`  // SessionID belongs to
	Father   *Agent       `json:"-"`        // If it's nil, it's the root agent of Manboster Session
	Children []*Agent     `json:"-"`        // The Agent's children
}

// Store stores Agent to DTO(Data Transfer Object)
func (a Agent) Store() Store {
	father := ""
	if a.Father != nil {
		father = a.Father.ID
	}
	var children []string
	if a.Children != nil && len(a.Children) > 0 {
		for _, child := range a.Children {
			children = append(children, child.ID)
		}
	}
	return Store{
		ID:       a.ID,
		Data:     a,
		Father:   father,
		Children: children,
	}
}
