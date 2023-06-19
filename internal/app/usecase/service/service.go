package service

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase struct {
	serviceRepo ServiceRepository
}

func NewUsecase(serviceRepo ServiceRepository) *Usecase {
	return &Usecase{
		serviceRepo: serviceRepo,
	}
}

func (u *Usecase) ServiceStatus(ctx context.Context) (*domain.StatusService, error) {
	return u.serviceRepo.GetCountTuples(ctx)
}

func (u *Usecase) ServiceClear(ctx context.Context) error {
	return u.serviceRepo.DeleteAllEntities(ctx)
}
