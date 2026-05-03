package file

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name     NameType `json:"name" description:"The name you want call, it would be enum, only 6 values: read, write, info, list, dir and delete. 'read' returns the content of the file, 'write' sets content to the file, 'list' shows file available in specific directory and 'delete' deletes file or folder specified. 'info' shows the information of specified file or folder. 'dir' shows the PWD of your session folder or the public folder(if is_public = 'true')" validate:"required" enum:"read,write,info,list,delete" example:"get"`
	FileName string   `json:"file_name" description:"The filename to use to read, write, get info or delete. If 'name' is 'list', it is not allowed to have any name." example:"example_file.md"`
	FilePath []string `json:"file_path" description:"The filepath to use, like ['dir1'] means this file is in the cmd directory, if there is no such filepath, we will create one for you." example:"['directory1']"`
	Content  string   `json:"content" description:"The content to use to write. Only valid when name='write'."`
	IsPublic bool     `json:"is_public" description:"Whether or not the file written is public or not, if this is true, it means your will be read/write in a session-shared public directory. If it's false, your data will be written or read in session-related file.'" example:"false"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
