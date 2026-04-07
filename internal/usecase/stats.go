package usecase

import (
	"context"

	"github.com/Puker228/user_echo/internal/domain"
	"github.com/Puker228/user_echo/internal/repository"
)

type StatsUseCase interface {
	Save(ctx context.Context, stats domain.UserStats) error
}

type statsUseCase struct {
	repo repository.StatsRepository
}

func NewStatsUseCase(repo repository.StatsRepository) StatsUseCase {
	return &statsUseCase{repo: repo}
}

func (uc *statsUseCase) Save(ctx context.Context, stats domain.UserStats) error {
	return uc.repo.Save(ctx, stats)
}
