package web_tracking

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// WebTrackingOption ...
type WebTrackingOption func(w *WebTracking)

// Reconfigure ...
func (w *WebTracking) Reconfigure(options ...WebTrackingOption) {
	for _, option := range options {
		option(w)
	}
}

// WithConfiguration ...
func WithConfiguration(config *WebTrackingConfig) WebTrackingOption {
	return func(w *WebTracking) {
		w.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) WebTrackingOption {
	return func(w *WebTracking) {
		w.logger = logger
		w.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) WebTrackingOption {
	return func(w *WebTracking) {
		w.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) WebTrackingOption {
	return func(w *WebTracking) {
		w.pm = mgr
	}
}
