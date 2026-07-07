package config

type ChannelType string

const (
	ChannelStable  = "stable"
	ChannelRC      = "rc"
	ChannelBeta    = "beta"
	ChannelAlpha   = "alpha"
	ChannelCanary  = "canary"
	ChannelNightly = "nightly"
	ChannelUnknown = "unknown"
)
