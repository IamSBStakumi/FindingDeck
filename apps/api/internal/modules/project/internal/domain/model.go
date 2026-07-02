package domain

import "time"

type Project struct {
	ID string
	Name string
	Description string
	RepositoryURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}