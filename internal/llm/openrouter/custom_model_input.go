package openrouter

import (
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/llm/oai_compat"
)

func (c *Config) InputCustomModel() (llm.Model, error) {
	customModel, err := oai_compat.InputModel()
	if err != nil {
		return llm.Model{}, err
	}

	for _, val := range c.inputModelData {
		if val == customModel.Name {
			return llm.Model{}, ErrDuplicatedModel
		}
	}
	return customModel, nil
}
