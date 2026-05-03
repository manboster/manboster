package browser

import (
	"encoding/json"
	"strings"
)

func (s *Service) CacheGroup(args string) string {
	arg := RunArgs{}
	var respStr strings.Builder
	if json.Unmarshal([]byte(args), &arg) == nil {
		respStr.WriteString(string(arg.Name))
	}
	// TODO: or domain based ...
	return respStr.String()
}
