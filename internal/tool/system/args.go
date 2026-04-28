package system

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name string `json:"name" description:"The name you want call, it would be enum, only 4 values: get, set, list and delete. 'get' returns the value of the key, 'set' sets value of the key to the database, 'list' shows the keys available in this database and 'delete' deletes the value from database." validate:"required" enum:"get,set,list,delete" example:"get"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
