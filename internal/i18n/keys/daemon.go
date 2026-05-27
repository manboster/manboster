package keys

// Daemon: execute
const (
	DaemonStartError       = "daemon.start_error"
	DaemonStartSuccess     = "daemon.start_success"
	DaemonStopSuccess      = "daemon.stop_success"
	DaemonStopStopped      = "daemon.stop_stopped"
	DaemonStopError        = "daemon.stop_error"
	DaemonRestartMessage   = "daemon.restart_message"
	DaemonStatusRunning    = "daemon.status_running"
	DaemonStatusNotRunning = "daemon.status_not_running"
	DaemonStatusError      = "daemon.status_error"
	DaemonNoConfig         = "daemon.no_config"
	DaemonLogError         = "daemon.log_error"
	DaemonLogReading       = "daemon.log_reading"
)
