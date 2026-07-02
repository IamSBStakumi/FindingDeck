package repository

import (
	"context"

	"github.com/IamSBStakumi/findingdeck/internal/modules/project/internal/domain"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *domain.Project) error
	List(ctx context.Context) ([]*domain.Project, error)
	FindByID(ctx context.Context, id string) (*domain.Project, error)
}