package postgresql

import (
	"context"
	"database/sql"

	"github.com/Puker228/user_echo/internal/domain"
)

type StatsRepository struct {
	db *sql.DB
}

func NewStatsRepository(db *sql.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (r *StatsRepository) Save(ctx context.Context, stats domain.UserStats) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO users (android_version, device_model, manufacturer, total_ram_gb, app_version)
		 VALUES ($1, $2, $3, $4, $5)`,
		stats.AndroidVersion, stats.DeviceModel, stats.Manufacturer, stats.TotalRamGB, stats.AppVersion,
	)
	return err
}
