package plugin

import (
	"github.com/manboster/manboster/spec/plugin"
)

type Provider interface {
	plugin.Provider
	Register() RegisterData
}

type RegisterData struct {
	Command    bool               // Required. can this plugin append to the command?
	Cron       bool               // Required. can this plugin use cron job?
	Schedule   string             // Required. schedule expression e.g. "0 * * * *"
	Sha512Sum  string             // Required. wasm file's sha512 sum
	Permission RegisterPermission // Required. Permission this plugin requested

	WebAccess  *RegisterWebAccessPayload  // web access payload
	FileSystem *RegisterFileSystemPayload // normally allow workspace, if you want to get more whitelist, pls use this to attach
	Network    *RegisterNetworkPayload    // network access payload
	Resource   *RegisterResourcePayload   // max resource
}
