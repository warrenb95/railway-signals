package repository

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/warrenb95/railway-signals/internal/domain"
)

// AddMileage addes a mileage value linking a signal with a track.
func (r *PostgresRepository) AddMileage(ctx context.Context, mileage *domain.Mileage) error {
	return r.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := r.db.Model(mileage).Table("mileages").Insert()
		if err != nil {
			r.logger.WithContext(ctx).WithError(err).Error("inserting signal mileage into store")
			return fmt.Errorf("inserting signal mileage: %w", err)
		}

		return nil
	})
}
