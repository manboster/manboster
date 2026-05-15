package cron

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name        NameType    `json:"name" description:"Operation to run: get, set, list, or delete scheduled jobs." validate:"required" enum:"get,set,list,delete" example:"get"`
	JobName     string      `json:"job_name" description:"Unique job name. Used by get, set, and delete." example:"daily_report"`
	MessageType MessageType `json:"message_type" description:"Message mode for set: text sends directly; prompt runs through the model." enum:"text,prompt" example:"prompt"`
	Prompt      string      `json:"prompt" description:"Message or prompt content. Required when name is set." example:"Send a brief daily report."`
	To          ToChatType  `json:"to" description:"Delivery target for set: this chat or the user's private chat." enum:"this,pm" example:"this"`
	Cron        string      `json:"cron" description:"Schedule expression. Use cron syntax or delay format such as +5m, +3h, or +7d." example:"0 8 * * *"`
	Ignore      IgnoreType  `json:"ignore" description:"Tool approval mode for prompt jobs: hachimi checks tool calls; ignore allows them automatically." example:"hachimi" enum:"hachimi,ignore"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
