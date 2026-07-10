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

type PackageType string

const (
	PackageGoInstall = "goinstall" // install via go install
	PackageAOSC      = "aosc"      // install via oma
	PackageAUR       = "aur"       // install via AUR
	PackageBrew      = "brew"      // install via Homebrew

	PackageUnknown = "unknown" // unknown source
	PackageNone    = ""
)
