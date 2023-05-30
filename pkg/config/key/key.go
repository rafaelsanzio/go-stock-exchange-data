package key

import (
	"github.com/joho/godotenv"

	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/applog"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/errs"
)

type Key struct {
	Name   string
	Secure bool
	Provider
}

type Provider string

var (
	ProviderStore  = Provider("store")
	ProviderEnvVar = Provider("env")
)

var (
	WSURL  = Key{Name: "WS_URL", Secure: false, Provider: ProviderEnvVar}
	WSPort = Key{Name: "WS_PORT", Secure: false, Provider: ProviderEnvVar}
	Topic  = Key{Name: "TOPIC", Secure: false, Provider: ProviderEnvVar}
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		_ = errs.ErrGettingEnv.Throwf(applog.Log, errs.ErrFmt, err)
	}
}
