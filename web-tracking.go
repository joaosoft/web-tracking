package web_tracking

import (
	"sync"

	"github.com/joaosoft/web"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type WebTracking struct {
	config        *WebTrackingConfig
	webClient     *web.Client
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	mux           sync.Mutex
}

// NewWebTracking ...
func NewWebTracking(options ...WebTrackingOption) (*WebTracking, error) {
	config, simpleConfig, err := NewConfig()

	webClient, err := web.NewClient()
	if err != nil {
		return nil, err
	}

	service := &WebTracking{
		pm:        manager.NewManager(manager.WithRunInBackground(false)),
		logger:    logger.NewLogDefault("web-tracking", logger.WarnLevel),
		config:    config.WebTracking,
		webClient: webClient,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(logger.Instance))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.WebTracking != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.WebTracking.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.WebTracking = &WebTrackingConfig{
			Host: defaultURL,
		}
	}

	service.Reconfigure(options...)

	web := service.pm.NewSimpleWebServer(config.WebTracking.Host)

	controller := NewController(service)
	controller.RegisterRoutes(web)

	service.pm.AddWeb("api_web_web_tracking", web)

	return service, nil
}

// Start ...
func (t *WebTracking) Start() error {
	return t.pm.Start()
}

// Stop ...
func (t *WebTracking) Stop() error {
	return t.pm.Stop()
}
