package cron

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name  string `json:"name" description:"The name you want call, it would be enum, only 4 values: get, set, list and delete. 'get' returns the value of the key, 'set' sets value of the key to the database, 'list' shows the keys available in this database and 'delete' deletes the value from database." validate:"required" enum:"get,set,list,delete" example:"get"`
	Key   string `json:"key" description:"The key to use to set, get or delete." example:"EXAMPLE_KEY"`
	Value string `json:"value" description:"The value to use to set, get or delete. On valid when 'name' is 'set'." example:"value"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
