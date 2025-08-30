package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chwiee/wb-semaphore/internal/domain"
)

type KeyService struct {
	BaseURL string
	Client  *http.Client
	Token   string
}

func NewKeyService(baseURL string, token string, client *http.Client) *KeyService {
	return &KeyService{BaseURL: baseURL, Client: client, Token: token}
}

func (s *KeyService) ListByProject(ctx context.Context, projectID int) ([]domain.Key, error) {
	u, _ := url.JoinPath(s.BaseURL, fmt.Sprintf("project/%d/keys", projectID))
	qs := url.Values{}
	qs.Set("Key type", "none")
	qs.Set("sort", "name")
	qs.Set("order", "asc")

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	req.Header.Set("Authorization", "Bearer "+s.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("keys request: %w", err)
	}
	defer resp.Body.Close()

	var out []domain.Key
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("keys decode: %w", err)
	}
	return out, nil
}
