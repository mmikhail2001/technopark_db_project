package service

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type ServiceRepository interface {
	GetCountTuples(ctx context.Context) (*domain.StatusService, error)
	DeleteAllEntities(ctx context.Context) error
}
