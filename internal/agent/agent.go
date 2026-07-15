package agent

import "github.com/manboster/manboster/spec/llm"

// Agent marks a new era of Manboster agenting
type Agent struct {
	powered  llm.Model
	events   []llm.Event
	cost     llm.Usage
	session  string
	father   *Agent
	children []*Agent
}
