package repository

import (
	"context"

	"github.com/Puker228/user_echo/internal/domain"
)

type StatsRepository interface {
	Save(ctx context.Context, stats domain.UserStats) error
}
