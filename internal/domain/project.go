package domain

import "time"

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Key struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Inventory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Repository struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID          int       `json:"id"`
	TemplateID  int       `json:"template_id"`
	ProjectID   int       `json:"project_id"`
	Status      string    `json:"status"`
	Playbook    string    `json:"playbook"`
	Secret      string    `json:"secret"`
	UserID      int       `json:"user_id"`
	Created     time.Time `json:"created"`
	InventoryID int       `json:"inventory_id"`
	Limit       string    `json:"limit"`
}

type Template struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
