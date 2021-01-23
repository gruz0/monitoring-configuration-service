package configuration

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) GetConfigurations(ctx context.Context) (configurations []Configuration, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetConfigurations", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetConfigurations(ctx)
}
