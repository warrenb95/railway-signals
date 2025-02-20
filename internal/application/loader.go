package application

import (
	"context"
	"fmt"

	"github.com/warrenb95/railway-signals/internal/domain"
)

// LoadTrackSignals stores the track signals.
func (a *Service) LoadTrackSignals(ctx context.Context, trackSignals []domain.TrackSignals) error {
	logger := a.Logger.WithContext(ctx)

	for _, ts := range trackSignals {
		err := a.TrackStore.CreateTrack(ctx, &domain.Track{
			ID:     ts.ID,
			Source: ts.Source,
			Target: ts.Target,
		})
		if err != nil {
			logger.WithError(err).Error("Failed to store track while loading track signals")
			return fmt.Errorf("creating track: %w", err)
		}

		for _, signal := range ts.Signals {
			// TODO: this needs to be removed when DB migration issues are fixed.
			if signal.ELR == "" {
				signal.ELR = "NULL"
			}
			err := a.SignalStore.CreateSignal(ctx, &domain.Signal{
				ID:   signal.ID,
				Name: signal.Name,
				ELR:  signal.ELR,
			})
			if err != nil {
				logger.WithError(err).Error("Failed to store signal while loading track signals")
				return fmt.Errorf("creating signal: %w", err)
			}

			err = a.MileageStore.AddMileage(ctx, &domain.Mileage{
				SignalID: signal.ID,
				TrackID:  ts.ID,
				Mileage:  signal.Mileage,
			})
			if err != nil {
				logger.WithError(err).Error("Failed to store signal mileage while loading track signals")
				return fmt.Errorf("creating mileage: %w", err)
			}
		}
	}
	return nil
}
