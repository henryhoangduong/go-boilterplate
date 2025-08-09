package logger

import (
	"fmt"
	"os"

	"github.com/henryhoangduong/go-boilterplate/internal/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type LoggerService struct {
	nrApp *newrelic.Application
}

func NewLoggerService(cfg *config.ObservabilityConfig) *LoggerService {
	service := &LoggerService{}
	if cfg.NewRelic.LicenseKey == "" {
		fmt.Println("New Relic license key not provided, skipping initialization")
		return service
	}
	var configOptions []newrelic.ConfigOption
	configOptions = append(configOptions,
		newrelic.ConfigAppName(cfg.ServiceName),
		newrelic.ConfigLicense(cfg.NewRelic.LicenseKey),
		newrelic.ConfigAppLogDecoratingEnabled(cfg.NewRelic.AppLogForwardingEnabled),
		newrelic.ConfigDistributedTracerEnabled(cfg.NewRelic.DistributedTracingEnabled),
	)

	if cfg.NewRelic.DebugLogging {
		configOptions = append(configOptions, newrelic.ConfigDebugLogger(os.Stdout))
	}

	app, err := newrelic.NewApplication(configOptions...)
	if err != nil {
		fmt.Printf("Failed to initialize New Relic: %v", err)
		return service
	}
	service.nrApp = app
	fmt.Printf("New relic initialized: %s", cfg.ServiceName)
	return service
}
