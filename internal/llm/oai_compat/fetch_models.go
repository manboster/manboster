package oai_compat

import "context"

// FetchModels gets models available.
func (s *Service) FetchModels(ctx context.Context) ([]string, error) {
	var modelInfos []string
	models, err := s.cli.ListModels(ctx)
	if err != nil {
		return []string{}, err
	}

	for _, m := range models.Models {
		modelInfos = append(modelInfos, m.ID)
	}

	return modelInfos, nil
}
