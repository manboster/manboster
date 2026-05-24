package file

import "github.com/manboster/manboster/spec/schema"

type ReadArgs struct {
	FileName string   `json:"file_name" description:"Target file name." example:"example.md" validate:"required"`
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]" validate:"required"`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false"`
}

type WriteArgs struct {
	FileName string   `json:"file_name" description:"Target file name." example:"example.md" validate:"required"`
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]" validate:"required"`
	Content  string   `json:"content" description:"Content to write to the file." validate:"required"`
	Append   bool     `json:"append" description:"If true, append content to the end of the file instead of overwriting it." example:"false"`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false"`
}

type InfoArgs struct {
	FileName string   `json:"file_name" description:"Target file name." example:"example.md" validate:"required"`
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]" validate:"required"`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false"`
}

type ListArgs struct {
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]" validate:"required"`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false" validate:"required"`
}

type DirArgs struct {
	IsPublic bool `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false" validate:"required"`
}

type DeleteArgs struct {
	FileName string   `json:"file_name" description:"Target file name." example:"example.md" validate:"required"`
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]" validate:"required"`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false" validate:"required"`
}

type GrepArgs struct {
	FileName string   `json:"file_name" description:"Target file name." example:"example.md" validate:"required"`
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]" validate:"required"`
	Keyword  string   `json:"keyword" description:"Keyword to search for in the file." example:"func main" validate:"required"`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false"`
}

type ReplaceArgs struct {
	FileName string   `json:"file_name" description:"Target file name." example:"example.md" validate:"required"`
	FilePath []string `json:"file_path" description:"Directory path segments under the workspace." example:"[\"notes\"]" validate:"required"`
	OldText  string   `json:"old_text" description:"The exact text to be replaced." validate:"required"`
	NewText  string   `json:"new_text" description:"The text to replace with." validate:"required"`
	Line     int      `json:"line" description:"Line number to restrict replacement to. Set to 0 to replace all occurrences in the file." example:"0"`
	IsPublic bool     `json:"is_public" description:"Use the shared public workspace instead of the current session workspace." example:"false"`
}

func (s *Service) Args() *schema.Args {
	return nil
}
