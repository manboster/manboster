package i18n

import (
	"embed"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

//go:embed locales/*
var localeFS embed.FS // embedded locales application
var localizer *i18n.Localizer
