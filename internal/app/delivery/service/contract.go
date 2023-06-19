package service

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase interface {
	ServiceStatus(ctx context.Context) (*domain.StatusService, error)
	ServiceClear(ctx context.Context) error
}
