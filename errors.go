package web_tracking

import (
	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

var (
	ErrorHandlingTracking = errors.New(errors.LevelError, int(web.StatusNotFound), "error handling tracking [error: %s]")
	ErrorTrackingStatus   = errors.New(errors.LevelError, 1, "error sending tracking [status: %d, error: %s]")
)
