package local

import (
	"os"

	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/config/key"
)

var Values = map[key.Key]string{
	key.WSURL:  getDefaultOrEnvVar("ws://websocket", "WS_URL"),
	key.WSPort: getDefaultOrEnvVar("8080", "WS_PORT"),
	key.Topic:  getDefaultOrEnvVar("stockexchange", "TOPIC"),
}

func getDefaultOrEnvVar(dfault, envVar string) string {
	val := os.Getenv(envVar)
	if val != "" {
		return val
	}
	return dfault
}
