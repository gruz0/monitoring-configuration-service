package configuration

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Configurations() (Configuration, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "list_configurations",
			"took", time.Since(begin),
		)
	}(time.Now())

	return s.Service.Configurations()
}
