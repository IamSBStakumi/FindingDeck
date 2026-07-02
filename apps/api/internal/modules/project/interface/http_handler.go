package project

import (
	"errors"
	"net/http"

	"github.com/IamSBStakumi/findingdeck/internal/generated/openapi"
	"github.com/IamSBStakumi/findingdeck/internal/modules/project/internal/domain"
	"github.com/IamSBStakumi/findingdeck/internal/modules/project/internal/usecase"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	service *usecase.Service
}

func NewHTTPHandler(service *usecase.Service) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (h *HTTPHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/projects", h.CreateProject)
	e.GET("/projects", h.ListProjects)
	e.GET("/projects/:projectID", h.GetProject)
}

func (h *HTTPHandler) CreateProject(c echo.Context) error {
	var request openapi.CreateProjectRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{
			Message: "Invalid request body",
		})
	}

	description := ""
	if request.Description != nil {
		description = *request.Description
	}

	project, err := h.service.Create(c.Request().Context(), usecase.CreateProjectInput{
		Name: request.Name,
		Description: description,
		RepositoryURL: request.RepositoryUrl,
	})

	if err != nil {
		if errors.Is(err, domain.ErrInvalidProjectInput) {
			return c.JSON(http.StatusBadRequest, openapi.ErrorResponse{
				Message: "Invalid project input",
			})
		}

		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Message: "Failed to create project",
		})
	}

	return c.JSON(http.StatusCreated, openapi.ProjectResponse{
		Project: toProjectDTO(*project),
	})
}

func (h *HTTPHandler) ListProjects(c echo.Context) error {
	projects, err := h.service.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Message: "Failed to list projects",
		})
	}

	return c.JSON(http.StatusOK, openapi.ListProjectResponse{
		Projects: toProjectDTOs(projects),
	})
}

func (h *HTTPHandler) GetProject(c echo.Context) error {
	projectID := c.Param("projectID")

	project, err := h.service.FindByID(c.Request().Context(), projectID)
	if err != nil {
		if errors.Is(err, domain.ErrProjectNotFound) {
			return c.JSON(http.StatusNotFound, openapi.ErrorResponse{
				Message: "Project not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Message: "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, openapi.ProjectResponse{
		Project: toProjectDTO(*project),
	})
}

func toProjectDTO(project domain.Project) openapi.Project {
	return openapi.Project{
		Id:            project.ID,
		Name:          project.Name,
		Description:   &project.Description,
		RepositoryUrl: project.RepositoryURL,
		CreatedAt:     project.CreatedAt,
		UpdatedAt:     project.UpdatedAt,
	}
}

func toProjectDTOs(projects []domain.Project) []openapi.Project {
	result := make([]openapi.Project, 0, len(projects))
	for _, project := range projects {
		result = append(result, toProjectDTO(project))
	}

	return result
}	