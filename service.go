package configuration

import (
	"context"
	"errors"
)

// Service is a simple CRUD interface for user profiles.
type Service interface {
	GetConfigurations(ctx context.Context) ([]Configuration, error)
}

// Configuration represents a single user profile.
// ID should be globally unique.
type Configuration struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type inmemService struct {
	m []Configuration
}

func NewInmemService() Service {
	return &inmemService{
		m: []Configuration{},
	}
}

func (s *inmemService) GetConfigurations(ctx context.Context) ([]Configuration, error) {
	return []Configuration{}, nil
}
