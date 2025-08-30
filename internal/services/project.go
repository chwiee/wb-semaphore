package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chwiee/wb-semaphore/internal/domain"
)

type ProjectService struct {
	BaseURL string
	Client  *http.Client
	Token   string
}

func NewProjectService(baseURL string, token string, client *http.Client) *ProjectService {
	return &ProjectService{BaseURL: baseURL, Client: client, Token: token}
}

func (s *ProjectService) List(ctx context.Context) ([]domain.Project, error) {
	u, _ := url.JoinPath(s.BaseURL, "projects")
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	req.Header.Set("Authorization", "Bearer "+s.Token)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("projects list request: %w", err)
	}
	defer resp.Body.Close()

	var out []domain.Project
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("projects decode: %w", err)
	}
	return out, nil
}

func (s *ProjectService) Get(ctx context.Context, id int) (*domain.Project, error) {
	u, _ := url.JoinPath(s.BaseURL, fmt.Sprintf("project/%d", id))

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Token)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("project get request: %w", err)
	}
	defer resp.Body.Close()

	var p domain.Project
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, fmt.Errorf("project decode: %w", err)
	}
	return &p, nil
}
