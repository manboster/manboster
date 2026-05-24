package file

type NameType string

const (
	NameRead    NameType = "read"
	NameWrite   NameType = "write"
	NameInfo    NameType = "info"
	NameDir     NameType = "dir"
	NameList    NameType = "list"
	NameDelete  NameType = "delete"
	NameGrep    NameType = "grep"
	NameReplace NameType = "replace"
)
