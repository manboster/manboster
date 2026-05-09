package cron

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name        NameType    `json:"name" description:"The name you want call, it would be enum, only 4 values: get, set, list and delete. 'get' returns the information of a job, 'set' sets the job to run, 'list' shows the jobs available in this database and 'delete' deletes the job from database." validate:"required" enum:"get,set,list,delete" example:"get"`
	JobName     string      `json:"job_name" description:"The job's name, it's unique. Only valid when 'name' = 'get' or 'set'" example:"job_cronjob"`
	MessageType MessageType `json:"message_type" description:"The message's type. Only valid when name = 'set'. It would be enum, only 2 values: 'text' and 'prompt', 'text' means the message would be sent as a text message, 'prompt' means the message would be user's prompt to activate the model." enum:"text,input" example:"text"`
	Prompt      string      `json:"prompt" description:"The message content. Only valid when 'name' = 'set'." example:"Please give a brief report to this user."`
	To          ToChatType  `json:"to" description:"Where the user will receive message. Only valid when 'name' = 'set'. It would be enum, only 2 values: 'this' and 'pm', 'this' means message would send in this chat, 'pm' means message would send in user's private chat" enum:"this,pm" example:"this"`
	Cron        string      `json:"cron" description:"Schedule for the task. Use cron expression (e.g. '0 8 * * *') for recurring tasks, or '+5m'/'+3h'/'+7d' format for one-time delayed execution." example:"* * * * *"`
	Ignore      IgnoreType  `json:"ignore" description:"Only valid when 'name' = 'set'. It would be enum, only 2 values: 'ignore' or 'hachimi', ignore means tool calls would be automatically executed while hachimi means hachimi enabled in tool call process and notify user" example:"hachimi" enum:"hachimi,ignore"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
