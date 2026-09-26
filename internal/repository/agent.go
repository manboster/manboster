package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
)

type AgentRepository interface {
	CreateAgent(ctx context.Context, agent types.Agent) error
	FindAgentByID(ctx context.Context, agentID string) (*types.Agent, error)
	UpdateAgent(ctx context.Context, agent types.Agent) error
	DeleteAgent(ctx context.Context, agentID string) error
}

type AgentRepo struct {
	db *gorm.DB
}

func NewAgentRepo(db *gorm.DB) *AgentRepo {
	return &AgentRepo{
		db: db,
	}
}

func (repo *AgentRepo) CreateAgent(ctx context.Context, agent types.Agent) error {
	agentDatabaseType := agent.Map()
	return repo.db.WithContext(ctx).Create(agentDatabaseType).Error
}

func (repo *AgentRepo) FindAgentByID(ctx context.Context, agentID string) (*types.Agent, error) {
	var agentInDB dbtypes.Agent
	res := repo.db.WithContext(ctx).First(&agentInDB, "agent_id = ?", agentID)
	if res.Error != nil {
		return nil, res.Error
	}
	agent, err := types.AgentFrom(agentInDB)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (repo *AgentRepo) UpdateAgent(ctx context.Context, agent types.Agent) error {
	agentDatabaseType := agent.Map()
	resp := repo.db.WithContext(ctx).Save(&agentDatabaseType)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (repo *AgentRepo) DeleteAgent(ctx context.Context, agentID string) error {
	resp := repo.db.WithContext(ctx).Where("agent_id = ?", agentID).Delete(&dbtypes.Agent{})
	if resp.RowsAffected == 0 {
		return ErrNotFound
	}
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
