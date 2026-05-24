package schema

// ToJSONSchema converts args to JSON Schema.
func (a *Args) ToJSONSchema() map[string]any {
	root := a

	if a == nil {
		return nil
	}

	properties := make(map[string]any)
	var required []string

	for _, prop := range root.Properties {
		paramDef := buildParamDef(prop)
		properties[prop.Name] = paramDef

		if prop.Required {
			required = append(required, prop.Name)
		}
	}

	res := map[string]any{
		"type":       "object",
		"properties": properties,
	}
	if len(required) > 0 {
		res["required"] = required
	}
	return res
}

func buildParamDef(arg *Args) map[string]any {
	def := map[string]any{
		"type":        arg.Type,
		"description": arg.Description,
	}

	if arg.IsEnum && len(arg.Enum) > 0 {
		def["enum"] = arg.Enum
	}

	if arg.Type == ArgsTypeObject && len(arg.Properties) > 0 {
		subProps := make(map[string]any)
		var subRequired []string
		for _, p := range arg.Properties {
			subProps[p.Name] = buildParamDef(p)
			if p.Required {
				subRequired = append(subRequired, p.Name)
			}
		}
		def["properties"] = subProps
		if len(subRequired) > 0 {
			def["required"] = subRequired
		}
	}

	if arg.Type == ArgsTypeArray && arg.Items != nil {
		def["items"] = buildParamDef(arg.Items)
	}

	return def
}
