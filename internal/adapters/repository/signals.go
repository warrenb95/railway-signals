package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/warrenb95/railway-signals/internal/domain"
)

// CreateSignal inserts a new signal into the database.
func (r *PostgresRepo) CreateSignal(ctx context.Context, signal *domain.Signal) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(signal).Table("signals").OnConflict("DO NOTHING").Insert()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting signal into store")
			return fmt.Errorf("inserting signal: %w", err)
		}

		return nil
	})
}

// GetSignal retrieves a signal by its ID.
func (r *PostgresRepo) GetSignal(ctx context.Context, signalID int) (*domain.Signal, error) {
	signal := &domain.Signal{ID: signalID}
	err := r.db.Model(signal).Table("signals").WherePK().Select()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("getting signal from store")
		return nil, fmt.Errorf("getting signal: %w", err)
	}

	return signal, nil
}

// ListSignals retrieves all signals from the database.
// Handles paginated requests and returns the total count along with the returned signals.
func (r *PostgresRepo) ListSignals(ctx context.Context, limit, page int) ([]domain.Signal, int, error) {
	var signals []domain.Signal

	err := r.db.Model(&signals).
		Limit(limit).
		Offset(page * limit). // TODO: validate limit and page
		Table("signals").
		Select()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("listing signals from store")
		return nil, 0, fmt.Errorf("listing signals: %w", err)
	}

	count, err := r.db.Model((*domain.Signal)(nil)).Count()
	if err != nil {
		r.logger.WithContext(ctx).WithError(err).Error("counting signals from store")
		return nil, 0, fmt.Errorf("counting signals: %w", err)
	}

	return signals, count, nil
}

// UpdateSignal modifies an existing signal.
func (r *PostgresRepo) UpdateSignal(ctx context.Context, updateReq *domain.Signal) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(updateReq).Table("signals").WherePK().Update()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("updating signal")
			return fmt.Errorf("updating signal: %w", err)
		}

		return nil
	})
}

// DeleteSignal removes a signal from the database.
func (r *PostgresRepo) DeleteSignal(ctx context.Context, signalID int) error {
	signal := &domain.Signal{ID: signalID}

	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(signal).Table("signals").WherePK().Delete()
		if err != nil && !errors.Is(err, pg.ErrNoRows) {
			r.logger.WithContext(ctx).WithError(err).Error("deleting signal")
			return fmt.Errorf("deleting signal: %w", err)
		}

		return err
	})
}
