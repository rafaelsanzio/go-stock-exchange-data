package config

import (
	"os"

	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/applog"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/config/key"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/config/local"

	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/errs"
)

// Default singleton pattern, similar to how Go does it in the log package
var (
	defaultService Service
)

// Load the default service
func init() {
	defaultService = local.Service{}
}

// Get a config value from the defaultService
func Value(k key.Key) (string, errs.AppError) {
	switch k.Provider {
	case key.ProviderStore:
		return defaultService.Value(k)
	case key.ProviderEnvVar:
		return os.Getenv(k.Name), nil
	default:
		return "", errs.ErrUnknownConfigProvider.Throw(applog.Log)
	}
}
