package file

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name     NameType `json:"name" description:"Operation to run: read, write, info, list, dir, or delete." validate:"required" enum:"read,write,info,list,delete" example:"read"`
	FileName string   `json:"file_name" description:"Target file name. Leave empty for list or dir." example:"example.md"`
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]"`
	Content  string   `json:"content" description:"Content to write. Used only when name is write."`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
