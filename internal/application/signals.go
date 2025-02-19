package application

import (
	"context"

	"github.com/warrenb95/railway-signals/internal/domain"
)

func (s *Service) CreateSignal(ctx context.Context, signal *domain.Signal) error {
	return s.SignalStore.CreateSignal(ctx, signal)
}

func (s *Service) GetSignal(ctx context.Context, signalID int) (*domain.Signal, error) {
	return s.SignalStore.GetSignal(ctx, signalID)
}

func (s *Service) ListSignals(ctx context.Context, limit, page int) (signals []domain.Signal, nextPage int, err error) {
	signals, count, err := s.SignalStore.ListSignals(ctx, limit, page)
	if err != nil {
		return nil, 0, err
	}

	if page*limit < count {
		nextPage = page + 1
	}

	return signals, nextPage, nil
}

func (s *Service) UpdateSignal(ctx context.Context, signal *domain.Signal) error {
	return s.SignalStore.UpdateSignal(ctx, signal)
}

func (s *Service) DeleteSignal(ctx context.Context, signalID int) error {
	return s.SignalStore.DeleteSignal(ctx, signalID)
}
