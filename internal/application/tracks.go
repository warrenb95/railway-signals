package application

import (
	"context"

	"github.com/warrenb95/railway-signals/internal/domain"
)

func (s *Service) GetSignalTracks(ctx context.Context, signalID, limit, page int) (tracks []domain.Track, nextPage int, err error) {
	// TODO: validate limit and page
	tracks, count, err := s.TrackStore.ListSignalTracks(ctx, signalID, limit, page)
	if err != nil {
		return nil, 0, err
	}

	if page*limit < count {
		nextPage = page + 1
	}

	return tracks, nextPage, nil
}

func (s *Service) CreateTrack(ctx context.Context, track *domain.Track) error {
	return s.TrackStore.CreateTrack(ctx, track)
}

func (s *Service) GetTrack(ctx context.Context, trackID int) (*domain.Track, error) {
	return s.TrackStore.GetTrack(ctx, trackID)
}

func (s *Service) ListTracks(ctx context.Context, limit, page int) (tracks []domain.Track, nextPage int, err error) {
	// TODO: validate limit and page
	signals, count, err := s.TrackStore.ListTracks(ctx, limit, page)
	if err != nil {
		return nil, 0, err
	}

	if page*limit < count {
		nextPage = page + 1
	}

	return signals, nextPage, nil
}

func (s *Service) UpdateTrack(ctx context.Context, track *domain.Track) error {
	return s.TrackStore.UpdateTrack(ctx, track)
}

func (s *Service) DeleteTrack(ctx context.Context, trackID int) error {
	return s.TrackStore.DeleteTrack(ctx, trackID)
}
