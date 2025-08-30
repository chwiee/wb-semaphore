package domain

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
	ID   int    `json:"id"`
	Name string `json:"name"`
}
