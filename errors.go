package web_tracking

import (
	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

var (
	ErrorHandlingTracking = errors.New(errors.ErrorLevel, int(web.StatusNotFound), "error handling tracking [error: %s]")
	ErrorTrackingStatus   = errors.New(errors.ErrorLevel, 1, "error sending tracking [status: %d, error: %s]")
)
