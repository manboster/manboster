package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

// Form holds the huh groups and a Collect function to retrieve filled values.
type Form struct {
	Groups []*huh.Group
	refs   []valueRef
	args   *Args
}

type valueRef struct {
	key string
	ptr any // *string, *bool, *[]string
}

// Collect returns all filled values as a nested map, suitable for mapstructure.
func (f *Form) Collect() map[string]any {
	result := make(map[string]any)
	for _, ref := range f.refs {
		var v any
		switch p := ref.ptr.(type) {
		case *string:
			v = *p
		case *bool:
			v = *p
		case *[]string:
			v = *p
		}
		setNested(result, ref.key, v)
	}
	return result
}

// Build rebuilds the form with the given initial values pre-filled.
// Clears any previous groups and refs.
func (f *Form) Build(config any) error {
	var values map[string]any

	switch v := config.(type) {
	case map[string]any:
		values = v
	case nil:
		// default vault
	default:
		// from JSON marshal to map
		data, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("build: marshal config: %w", err)
		}
		if err := json.Unmarshal(data, &values); err != nil {
			return fmt.Errorf("build: unmarshal config: %w", err)
		}
	}

	f.Groups = f.Groups[:0]
	f.refs = f.refs[:0]
	collectGroups(f.args.Nodes, &f.Groups, &f.refs, "", values)
	return nil
}

func setNested(m map[string]any, key string, val any) {
	parts := strings.Split(key, ".")
	for i := 0; i < len(parts)-1; i++ {
		existing, ok := m[parts[i]]
		if ok {
			if nested, ok := existing.(map[string]any); ok {
				m = nested
				continue
			}
		}
		newMap := make(map[string]any)
		m[parts[i]] = newMap
		m = newMap
	}
	m[parts[len(parts)-1]] = val
}
