package request

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Verb    string `json:"verb" description:"HTTP Verb to use" validate:"required" example:"GET" enum:"GET,POST,PUT,HEAD,OPTIONS,DELETE,PATCH,TRACE"`
	Timeout int    `json:"timeout" description:"Maximum execution time in seconds. Default is 120." validate:"required" example:"120"`
	Headers string `json:"headers" description:"HTTP Headers to use in request body" example:"{\"Authorization\":\"Bearer xxxxxxxxxxxxxxxx\"}"`
	Payload string `json:"payload" description:"HTTP Payload to use in request body" example:"payload"`
	URL     string `json:"url" description:"HTTP URL to use in request body" example:"https://example.com"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
