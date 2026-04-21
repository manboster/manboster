package config

type VersionType int8

const (
	VersionStable VersionType = iota
	VersionRC
	VersionBeta
	VersionAlpha
	VersionCanary
)

func (v VersionType) String() string {
	switch v {
	case VersionStable:
		return "stable"
	case VersionRC:
		return "release candidate"
	case VersionBeta:
		return "beta"
	case VersionAlpha:
		return "alpha"
	case VersionCanary:
		return "canary"
	default:
		return "unknown"
	}
}
