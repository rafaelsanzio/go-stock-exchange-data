package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"

	"github.com/rafaelsanzio/go-stock-exchange-data/data"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/applog"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/config"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/config/key"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/errs"
)

func Run(ctx context.Context) {
	events := data.ReadJSONData()

	wsURL, err_ := config.Value(key.WSURL)
	if err_ != nil {
		_ = errs.ErrGettingEnvWebSocketURL.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	wsPort, err_ := config.Value(key.WSPort)
	if err_ != nil {
		_ = errs.ErrGettingEnvWebSocketPort.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	topic, err_ := config.Value(key.Topic)
	if err_ != nil {
		_ = errs.ErrGettingEnvTopic.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	wsListenerURL := fmt.Sprintf("%s:%s/ws?topic=%s", wsURL, wsPort, topic)

	conn, _, err := websocket.DefaultDialer.Dial(wsListenerURL, nil)
	if err != nil {
		_ = errs.ErrWebSocketConnection.Throwf(applog.Log, errs.ErrFmt, err)
	}
	defer conn.Close()

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		event := events[rand.Intn(len(events))]

		eventJSON, err := json.Marshal(event)
		if err != nil {
			_ = errs.ErrMarshalingJson.Throwf(applog.Log, errs.ErrFmt, err)
			continue
		}

		err = conn.WriteMessage(websocket.TextMessage, eventJSON)
		if err != nil {
			_ = errs.ErrWebSocketSendingMsg.Throwf(applog.Log, errs.ErrFmt, err)
			continue
		}

		log.Printf("Sent message to websocket: %s\n", string(eventJSON))
	}
}
