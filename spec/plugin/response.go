package plugin

type RunResponse struct {
	Response string `json:"response"` // This response is for LLM.

	Hangup     bool              `json:"hangup"`               // is this hang up or not?
	SessionId  string            `json:"sessionId,omitempty"`  // session id in the stateful tool
	Operations []EngineOperation `json:"operations,omitempty"` // Operations wants to run
}
