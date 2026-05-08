package gguf

import (
	"regexp"
	"strconv"

	"github.com/manboster/manboster/internal/hachimi"
)

var hachimiRegex = regexp.MustCompile(`Safety:\s*(Safe|Unsafe|Controversial)\s*\n?\s*Categories:\s*(.*)`)

func (s *Service) purgeSafeguardChatData(resp string) (*hachimi.Response, error) {
	res := &hachimi.Response{}

	matches := hachimiRegex.FindStringSubmatch(resp)
	if len(matches) < 2 {
		res.Type = hachimi.ResponseStatusUnsafe
		res.Reason = "Model output error, could not complie regex!"
		return res, nil
	}
	switch matches[1] {
	case "Safe":
		res.Type = hachimi.ResponseStatusSafe
	case "Unsafe":
		res.Type = hachimi.ResponseStatusUnsafe
	case "Controversial":
		res.Type = hachimi.ResponseStatusInspect
	}
	if len(matches) > 2 {
		res.Reason = matches[2]
	}

	return res, nil
}

func (s *Service) purgeLLMChatData(resp string) (*hachimi.Response, error) {
	num, err := strconv.Atoi(resp)
	if err != nil {
		return nil, err
	}

	res := &hachimi.Response{}

	if num >= 80 {
		res.Type = hachimi.ResponseStatusUnsafe
	} else if num >= 60 {
		res.Type = hachimi.ResponseStatusInspect
	} else {
		res.Type = hachimi.ResponseStatusSafe
	}
	res.Reason = ""
	return res, nil
}
