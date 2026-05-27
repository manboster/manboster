package keys

// Daemon: execute
const (
	AppDaemonStartError       = "daemon.start_error"
	AppDaemonStartSuccess     = "daemon.start_success"
	AppDaemonStopSuccess      = "daemon.stop_success"
	AppDaemonStopStopped      = "daemon.stop_stopped"
	AppDaemonStopError        = "daemon.stop_error"
	AppDaemonRestartMessage   = "daemon.restart_message"
	AppDaemonStatusRunning    = "daemon.status_running"
	AppDaemonStatusNotRunning = "daemon.status_not_running"
	AppDaemonStatusError      = "daemon.status_error"
	AppDaemonNoConfig         = "daemon.no_config"
	AppDaemonLogError         = "daemon.log_error"
	AppDaemonLogReading       = "daemon.log_reading"
)
