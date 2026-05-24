package cron

import "github.com/manboster/manboster/spec/schema"

type SetArgs struct {
	JobName     string      `json:"job_name" description:"Unique job name." example:"daily_report" validate:"required"`
	MessageType MessageType `json:"message_type" description:"Message mode: text sends directly; prompt runs through the model." enum:"text,prompt" example:"prompt" validate:"required"`
	Prompt      string      `json:"prompt" description:"Message or prompt content." example:"Send a brief daily report." validate:"required"`
	To          ToChatType  `json:"to" description:"Delivery target: this chat or the user's private chat." enum:"this,pm" example:"this" validate:"required"`
	Cron        string      `json:"cron" description:"Schedule expression. Use 6-field cron with seconds (sec min hour dom mon dow), or delay format such as +5m, +3h, or +7d." example:"0 0 8 * * *" validate:"required"`
	Ignore      IgnoreType  `json:"ignore" description:"Tool approval mode: hachimi checks tool calls; ignore allows them automatically." example:"hachimi" enum:"hachimi,ignore"`
}

type GetArgs struct {
	JobName string `json:"job_name" description:"Job name to retrieve." example:"daily_report" validate:"required"`
}

type DeleteArgs struct {
	JobName string `json:"job_name" description:"Job name to delete." example:"daily_report" validate:"required"`
}

func (s *Service) Args() *schema.Args {
	return nil
}
