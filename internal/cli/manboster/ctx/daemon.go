package ctx

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/sevlyar/go-daemon"
)

var DaemonCtx = &daemon.Context{
	PidFileName: config.Path("manboster.pid"),
	PidFilePerm: 0644,
	LogFileName: config.Path("manboster.log"),
	LogFilePerm: 0640,
	WorkDir:     config.Path(""),
	Umask:       027,
}
