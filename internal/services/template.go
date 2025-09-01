package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/chwiee/wb-semaphore/internal/domain"
)

type TemplateService struct {
	BaseURL string
	Client  *http.Client
	Token   string
}

func NewTemplateService(baseURL string, token string, client *http.Client) *TemplateService {
	return &TemplateService{BaseURL: baseURL, Client: client, Token: token}
}

func (s *TemplateService) ListByProject(ctx context.Context, projectID int) ([]domain.Template, error) {
	u, _ := url.JoinPath(s.BaseURL, fmt.Sprintf("project/%d/templates", projectID))
	qs := url.Values{}
	qs.Set("sort", "name")
	qs.Set("order", "asc")

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u+"?"+qs.Encode(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Token)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tempalte request: %w", err)
	}
	defer resp.Body.Close()

	var out []domain.Template
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("templates decode: %w", err)
	}
	return out, nil
}

func (s *TemplateService) Filter(ctx context.Context, projectID int, template string) (*domain.Template, error) {
	templates, err := s.ListByProject(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("templates lists reqeust: %w", err)
	}

	for _, t := range templates {
		if strings.EqualFold(t.Name, template) {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("template '%s' not found", template)
}
