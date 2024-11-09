package web

import "time"

const (
	TimeoutShutdown = 30 * time.Second
)

const (
	LogFormatStartShutdown  = "Starting shutdown process, timeout: %s"
	LogFormatFailedShutdown = "Failed to shutdown: %+v"
)

const (
	LogMessageForceShutdown                  = "Force shutdown"
	LogMessageShutdownCompleted              = "Graceful shutdown completed"
	LogMessageRepositoryProcessingCompleted  = "Repository processing completed"
	LogMessageApplicationProcessingCompleted = "Application processing completed"
)

const (
	HeaderAuthorization = "Authorization"
)

const (
	PrefixAuthorizationToken = "Bearer "
)
