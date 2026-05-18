package config

import (
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/spec/cli"
)

// CliForm holds collected values filled via a cli.Provider.
// It mirrors the Form API but drives interaction through cli.Provider
// instead of huh groups.
type CliForm struct {
	values map[string]any
	args   *Args
}

// Collect returns all filled values as a nested map, suitable for mapstructure.
func (f *CliForm) Collect() map[string]any {
	return f.values
}

// Build re-runs the provider interaction with the given initial values pre-filled.
// Clears any previous values.
func (f *CliForm) Build(p cli.Provider, config any) error {
	var initial map[string]any

	switch v := config.(type) {
	case map[string]any:
		initial = v
	case nil:
		// no defaults
	default:
		data, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("build: marshal config: %w", err)
		}
		if err := json.Unmarshal(data, &initial); err != nil {
			return fmt.Errorf("build: unmarshal config: %w", err)
		}
	}

	f.values = make(map[string]any)
	return collectProviderValues(f.args.Nodes, p, f.values, "", initial)
}
