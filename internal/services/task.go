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

type TaskService struct {
	BaseURL string
	Client  *http.Client
	Token   string
}

func NewTaskService(baseURL string, token string, client *http.Client) *TaskService {
	return &TaskService{BaseURL: baseURL, Client: client, Token: token}
}

func (s *TaskService) Run(ctx context.Context, projectID int, inventoryID int, templateID int) (*domain.Task, error) {
	u, _ := url.JoinPath(s.BaseURL, fmt.Sprintf("project/%d/tasks", projectID))
	body := map[string]interface{}{
		"inventory_id": inventoryID,
		"template_id":  templateID,
	}

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, u, strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Token)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tasks request: %w", err)
	}
	defer resp.Body.Close()

	var out domain.Task
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("tasks decode: %w", err)
	}
	return &out, nil
}
