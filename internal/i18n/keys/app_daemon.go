package keys

// Daemon: execute
const (
	AppDaemonStartError       = "app.daemon.start_error"
	AppDaemonStartSuccess     = "app.daemon.start_success"
	AppDaemonStopSuccess      = "app.daemon.stop_success"
	AppDaemonStopStopped      = "app.daemon.stop_stopped"
	AppDaemonStopError        = "app.daemon.stop_error"
	AppDaemonRestartMessage   = "app.daemon.restart_message"
	AppDaemonStatusRunning    = "app.daemon.status_running"
	AppDaemonStatusNotRunning = "app.daemon.status_not_running"
	AppDaemonStatusError      = "app.daemon.status_error"
	AppDaemonNoConfig         = "app.daemon.no_config"
	AppDaemonLogError         = "app.daemon.log_error"
	AppDaemonLogReading       = "app.daemon.log_reading"
)
