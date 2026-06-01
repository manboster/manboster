package skill

func Load(path string, name string, isDir bool) error {
	skill, err := LoadSkill(path, name, isDir)
	if err != nil {
		return err
	}
	Register(name, skill)
	return nil
}
