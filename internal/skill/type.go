package skill

// Frontmatter is used to handle and read frontmatter type in the skills
type Frontmatter struct {
	Name            string `yaml:"name"`
	Description     string `yaml:"description"`
	Homepage        string `yaml:"homepage,omitempty"`
	UserInvocable   bool   `yaml:"user-invocable,omitempty"`
	Hidden          bool   `yaml:"hidden,omitempty"`
	CommandDispatch string `yaml:"command-dispatch,omitempty"`
	CommandName     string `yaml:"command-name,omitempty"`
	// read raw metadata message first
	MetadataRaw string `yaml:"metadata,omitempty"`
}

type Metadata struct {
	OpenClaw OpenClawMeta `json:"openclaw"`
}

// OpenClawMeta is the core compatibility layer for manboster and openclaw
type OpenClawMeta struct {
	Requires   Requires       `json:"requires,omitempty"`
	PrimaryEnv string         `json:"primaryEnv,omitempty"`
	Emoji      string         `json:"emoji,omitempty"`
	Homepage   string         `json:"homepage,omitempty"`
	Platforms  []string       `json:"platforms,omitempty"`
	SkillKey   string         `json:"skillKey,omitempty"`
	Install    []InstallEntry `json:"install,omitempty"`
}

type Requires struct {
	Bins   []string `json:"bins,omitempty"`
	Env    []string `json:"env,omitempty"`
	Config []string `json:"config,omitempty"`
}

type InstallEntry struct {
	ID      string   `json:"id"`
	Kind    string   `json:"kind"`
	Formula string   `json:"formula,omitempty"`
	Bins    []string `json:"bins,omitempty"`
	Label   string   `json:"label,omitempty"`
}

// RawData is purged from reading
type RawData struct {
	PWD   string      // Present Work Dir
	Front Frontmatter // Frontmatter YAML raw data
	Meta  Metadata    // parsed data
	Body  string      // the body content
}
