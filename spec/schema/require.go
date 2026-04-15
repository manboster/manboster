package schema

type RequirementData struct {
	Type       RequirementType
	Optional   bool
	MinVersion int    // for simplicity, we only check major version
	Package    string // Package name, i.e. dev.manboster.github(official package) or weekly-report(skills) or maimai-b50-checker(skills)
}

type RequirementType int16

const (
	RequireExecutable RequirementType = iota // executable file
	RequireTool                              // first find in tool, then in plugin
	RequireSkill                             // skills required to get
)
