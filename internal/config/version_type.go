package config

type VersionType string

const (
	VersionStable  = "stable"
	VersionRC      = "rc"
	VersionBeta    = "beta"
	VersionAlpha   = "alpha"
	VersionCanary  = "canary"
	VersionUnknown = "unknown"
)
