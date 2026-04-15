package plugin

type RegisterPermission int

const (
	RegisterPermissionFileSystem RegisterPermission = 1 << iota
	RegisterPermissionNetwork
	RegisterPermissionWebAccess
	RegisterPermissionExecuteCommand
	RegisterPermissionGetScreenshot
	RegisterPermissionKeyboardInput
	RegisterPermissionSimulateTouch
	RegisterPermissionOccupyMoreResources
)

func (r RegisterPermission) Has(permission RegisterPermission) bool {
	return r&permission == permission
}
