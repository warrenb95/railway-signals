package application

import (
	"github.com/sirupsen/logrus"
	"github.com/warrenb95/railway-signals/internal/domain"
)

type Service struct {
	Logger *logrus.Logger

	SignalStore  domain.SignalStore
	TrackStore   domain.TrackStore
	MileageStore domain.MileageStore
}
