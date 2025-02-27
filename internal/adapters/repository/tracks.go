package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/warrenb95/railway-signals/internal/domain"
)

// CreateTrack inserts a new track into the database.
func (r *PostgresRepository) CreateTrack(ctx context.Context, track *domain.Track) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.ModelContext(ctx, track).OnConflict("DO NOTHING").Insert()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting track into store")
			return fmt.Errorf("inserting track: %w", err)
		}
		return nil
	})
}

// GetTrack retrieves a track by its ID.
func (r *PostgresRepository) GetTrack(ctx context.Context, trackID int) (*domain.Track, error) {
	track := &domain.Track{ID: trackID}

	err := r.db.ModelContext(ctx, track).WherePK().Select()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("getting track from store")
		return nil, fmt.Errorf("getting track: %w", err)
	}

	return track, nil
}

// ListTracks retrieves all tracks from the database.
// Handles paginated requests and returns the total count along with the returned tracks.
func (r *PostgresRepository) ListTracks(ctx context.Context, limit, page int) ([]domain.Track, int, error) {
	var tracks []domain.Track
	count, err := r.db.ModelContext(ctx, &tracks).
		Limit(limit).
		Offset(page * limit).
		SelectAndCount()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("listing tracks from store")
		return nil, 0, fmt.Errorf("listing signals: %w", err)
	}

	return tracks, count, nil
}

// UpdateTrack modifies an existing track.
func (r *PostgresRepository) UpdateTrack(ctx context.Context, track *domain.Track) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.ModelContext(ctx, track).WherePK().Update()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("updating track")
			return fmt.Errorf("updating track: %w", err)
		}

		return nil
	})
}

// DeleteTrack removes a track from the database.
func (r *PostgresRepository) DeleteTrack(ctx context.Context, trackID int) error {
	track := &domain.Track{ID: trackID}
	_, err := r.db.ModelContext(ctx, track).Table("tracks").WherePK().Delete()
	if err != nil && !errors.Is(err, pg.ErrNoRows) {
		r.logger.WithContext(ctx).WithError(err).Error("deleting track")
		return fmt.Errorf("deleting track: %w", err)
	}

	return err
}

// ListSignalTracks retrieves all tracks from the database associated with the given signal.
// Handles paginated requests and returns the total count along with the returned tracks.
func (r *PostgresRepository) ListSignalTracks(ctx context.Context, signalID, limit, page int) ([]domain.Track, int, error) {
	var tracks []domain.Track
	err := r.db.ModelContext(ctx, &tracks).
		Limit(limit).
		Offset(page*limit).
		Join("mileage ON tracks.id=mileage.track_id").
		Where("mileage.signal_id = ?", signalID).
		Select()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("listing tracks from store")
		return nil, 0, fmt.Errorf("listing signals: %w", err)
	}

	count, err := r.db.ModelContext(ctx, (*domain.Track)(nil)).Table("tracks").Count()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("counting tracks from store")
		return nil, 0, fmt.Errorf("counting tracks: %w", err)
	}

	return tracks, count, nil
}
