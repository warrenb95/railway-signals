package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/warrenb95/railway-signals/internal/domain"
)

// CreateSignal inserts a new signal into the database.
func (r *PostgresRepository) CreateSignal(ctx context.Context, signal *domain.Signal) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(signal).OnConflict("DO NOTHING").Insert()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting signal into store")
			return fmt.Errorf("inserting signal: %w", err)
		}

		return nil
	})
}

// GetSignal retrieves a signal by its ID.
func (r *PostgresRepository) GetSignal(ctx context.Context, signalID int) (*domain.Signal, error) {
	signal := &domain.Signal{ID: signalID}
	err := r.db.Model(signal).WherePK().Select()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("getting signal from store")
		return nil, fmt.Errorf("getting signal: %w", err)
	}

	return signal, nil
}

// ListSignals retrieves all signals from the database.
// Handles paginated requests and returns the total count along with the returned signals.
func (r *PostgresRepository) ListSignals(ctx context.Context, limit, page int) ([]domain.Signal, int, error) {
	var signals []domain.Signal

	count, err := r.db.Model(&signals).
		Limit(limit).
		Offset(page * limit).
		SelectAndCount()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("listing signals from store")
		return nil, 0, fmt.Errorf("listing signals: %w", err)
	}

	return signals, count, nil
}

// UpdateSignal modifies an existing signal.
func (r *PostgresRepository) UpdateSignal(ctx context.Context, updateReq *domain.Signal) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(updateReq).WherePK().Update()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("updating signal")
			return fmt.Errorf("updating signal: %w", err)
		}

		return nil
	})
}

// DeleteSignal removes a signal from the database.
func (r *PostgresRepository) DeleteSignal(ctx context.Context, signalID int) error {
	signal := &domain.Signal{ID: signalID}

	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(signal).WherePK().Delete()
		if err != nil && !errors.Is(err, pg.ErrNoRows) {
			r.logger.WithContext(ctx).WithError(err).Error("deleting signal")
			return fmt.Errorf("deleting signal: %w", err)
		}

		return err
	})
}
