package web_tracking

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	WebTracking *WebTrackingConfig `json:"web-tracking"`
}

// WebTrackingConfig ...
type WebTrackingConfig struct {
	Host         string `json:"host"`
	TrackingHost string `json:"tracking_host"`
	Log          struct {
		Level string `json:"level"`
	} `json:"log"`
}

// newConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", getEnv()), appConfig)

	return appConfig, simpleConfig, err
}
