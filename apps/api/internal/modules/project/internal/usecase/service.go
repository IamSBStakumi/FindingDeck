package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/IamSBStakumi/findingdeck/internal/modules/project/internal/domain"
	"github.com/IamSBStakumi/findingdeck/internal/modules/project/internal/domain/repository"
	"github.com/google/uuid"
)

var ErrInvalidProjectInput = errors.New("invalid project input")

type Service struct {
	repository repository.ProjectRepository
}

func NewService(repository repository.ProjectRepository) *Service {
	return &Service{
		repository: repository,
	}
}

type CreateProjectInput struct {
	Name          string
	Description   string
	RepositoryURL string
}

func (s *Service) Create(ctx context.Context, input CreateProjectInput) (*domain.Project, error) {
	name := strings.TrimSpace(input.Name)
	repositoryURL := strings.TrimSpace(input.RepositoryURL)

	if name == "" || repositoryURL == "" {
		return nil, ErrInvalidProjectInput
	}

	now := time.Now().UTC()

	project := domain.Project{
		ID:            newProjectID(),
		Name:          name,
		Description:   strings.TrimSpace(input.Description),
		RepositoryURL: repositoryURL,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := s.repository.Create(ctx, project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *Service) List(ctx context.Context) ([]domain.Project, error) {
	return s.repository.List(ctx)
}

func (s *Service) FindByID(ctx context.Context, id string) (*domain.Project, error) {
	return s.repository.FindByID(ctx, id)
}

func newProjectID() string {
	// Generate a new UUID for the project ID
	return uuid.New().String()
}