package types

import (
	"encoding/json"

	"github.com/manboster/manboster/internal/database/types"
)

// Agent stores agents in Manboster
type Agent struct {
	ID       string
	Data     string
	Father   string
	Children []string
}

func (agent Agent) Map() types.Agent {
	data, _ := json.Marshal(agent.Children)
	return types.Agent{
		AgentID:  agent.ID,
		Data:     agent.Data,
		Father:   agent.Father,
		Children: string(data),
	}
}

func AgentFrom(agent types.Agent) (Agent, error) {
	var children []string
	err := json.Unmarshal([]byte(agent.Children), &children)
	if err != nil {
		return Agent{}, err
	}

	return Agent{
		ID:       agent.AgentID,
		Data:     agent.Data,
		Children: children,
		Father:   agent.Father,
	}, nil
}
