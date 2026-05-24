package skill

// Skill defines OpenClaw skill compact API for Manboster, use it to read, parse and inject into.
type Skill struct {
	Name           string // skill name, read by file name, but `~/.manboster/skills/a.md` or `~/.manboster/skills/a/` is seen as the same so they can't be duplicated.
	IsDirectory    bool   // is this skill a directory or not?
	Description    string // this skill's description
	DisplayName    string // display for human, read from SKILLS.md
	RepresentEmoji string // An emoji can represent for the skill, the default value is "📝".
}
