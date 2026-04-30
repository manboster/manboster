package config

import (
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
func (f *Form) Build(initialValues map[string]any) error {
	f.Groups = make([]*huh.Group, 0)
	f.refs = make([]valueRef, 0)
	// We need access to args here. If Form stores a reference to the source Args,
	// use it; otherwise pass it in. Assuming Form holds the original args pointer:
	if f.args == nil {
		return nil
	}
	collectGroups(f.args.Nodes, &f.Groups, &f.refs, "", initialValues)
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
