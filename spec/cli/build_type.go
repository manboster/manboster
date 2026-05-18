package cli

type BuildableType struct {
	name        string
	displayName string
}

func NewBuildableType(name string, displayName string) *BuildableType {
	return &BuildableType{
		name:        name,
		displayName: displayName,
	}
}

func (b *BuildableType) Name() string {
	return b.name
}

func (b *BuildableType) DisplayName() string {
	return b.displayName
}
