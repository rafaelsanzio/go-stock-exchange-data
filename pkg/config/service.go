package config

import (
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/config/key"
	"github.com/rafaelsanzio/go-stock-exchange-data/pkg/errs"
)

type Service interface {
	Value(key.Key) (string, errs.AppError)
}
