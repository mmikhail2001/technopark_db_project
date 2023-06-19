package vote

import (
	"context"

	"github.com/mmikhail2001/technopark_db_project/internal/app/domain"
)

type Usecase interface {
	Vote(ctx context.Context, vote *domain.Vote, thread *domain.Thread) (*domain.Thread, error)
}
