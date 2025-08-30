package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chwiee/wb-semaphore/internal/domain"
)

type InventoryService struct {
	BaseURL string
	Client  *http.Client
	Token   string
}

func NewInventoryService(baseURL string, token string, client *http.Client) *InventoryService {
	return &InventoryService{BaseURL: baseURL, Client: client, Token: token}
}

func (s *InventoryService) ListByProject(ctx context.Context, projectID int) ([]domain.Inventory, error) {
	u, _ := url.JoinPath(s.BaseURL, fmt.Sprintf("project/%d/inventory", projectID))
	qs := url.Values{}
	qs.Set("sort", "name")
	qs.Set("order", "asc")

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u+"?"+qs.Encode(), nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Token)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("inventory request: %w", err)
	}
	defer resp.Body.Close()

	var out []domain.Inventory
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("inventory decode: %w", err)
	}
	return out, nil
}
