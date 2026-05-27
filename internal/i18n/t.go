package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func T(messageID string, args ...any) string {
	if localizer == nil {
		err := Init()
		if err != nil {
			return ""
		}
		return messageID
	}

	config := &i18n.LocalizeConfig{
		MessageID: messageID,
	}

	for _, arg := range args {
		switch v := arg.(type) {
		case map[string]any:
			config.TemplateData = v
		case int:
			config.PluralCount = v
			if config.TemplateData == nil {
				config.TemplateData = map[string]any{"Count": v}
			}
		}
	}

	msg, err := localizer.Localize(config)
	if err != nil {
		return messageID
	}

	return msg
}

// Te -> Translate Easy Mode, only passthrough name and error!
func Te(messageID string, name string, err error) string {
	arr := map[string]any{}
	if err != nil {
		arr["Error"] = err.Error()
	}
	if name != "" {
		arr["Name"] = name
	}
	return T(messageID, arr)
}
