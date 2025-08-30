package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chwiee/wb-semaphore/internal/config"
	"github.com/chwiee/wb-semaphore/internal/services"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

type ProjectHandler struct {
	Projects  *services.ProjectService
	Keys      *services.KeyService
	Inventory *services.InventoryService
	Repos     *services.RepoService
	Tasks     *services.TaskService
}

func NewProjectHandler() *ProjectHandler {
	base := config.GetSemaphoreURL()
	token := config.GetBearerToken()
	client := services.NewHTTPClient()

	return &ProjectHandler{
		Projects:  services.NewProjectService(base, token, client),
		Keys:      services.NewKeyService(base, token, client),
		Inventory: services.NewInventoryService(base, token, client),
		Repos:     services.NewRepoService(base, token, client),
		Tasks:     services.NewTaskService(base, token, client),
	}
}

// GET /projects  -> "ID: X - Project: NAME"
func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	list, err := h.Projects.List(ctx)
	if err != nil {
		http.Error(w, "failed to list projects: "+err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for _, p := range list {
		fmt.Fprintf(w, "ID: %d - Project: %s\n", p.ID, p.Name)
	}
}

// GET /project/{id}
func (h *ProjectHandler) GetProjectDetail(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	projectID, err := strconv.Atoi(idStr)
	if err != nil || projectID <= 0 {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	g, gctx := errgroup.WithContext(ctx)

	var (
		projName string
		keysOut  = make([]string, 0)
		invOut   = make([]string, 0)
		repOut   = make([]string, 0)
		taskOut  = make([]string, 0)
	)

	g.Go(func() error {
		p, err := h.Projects.Get(gctx, projectID)
		if err != nil {
			return err
		}
		projName = p.Name
		return nil
	})

	g.Go(func() error {
		ks, err := h.Keys.ListByProject(gctx, projectID)
		if err != nil {
			return err
		}
		for _, k := range ks {
			keysOut = append(keysOut, fmt.Sprintf("ID: %d - %s", k.ID, k.Name))
		}
		return nil
	})

	g.Go(func() error {
		inv, err := h.Inventory.ListByProject(gctx, projectID)
		if err != nil {
			return err
		}
		for _, v := range inv {
			invOut = append(invOut, fmt.Sprintf("ID: %d - %s", v.ID, v.Name))
		}
		return nil
	})

	g.Go(func() error {
		repos, err := h.Repos.ListByProject(gctx, projectID)
		if err != nil {
			return err
		}
		for _, r := range repos {
			repOut = append(repOut, fmt.Sprintf("ID: %d - %s", r.ID, r.Name))
		}
		return nil
	})

	g.Go(func() error {
		tasks, err := h.Tasks.ListByProject(gctx, projectID)
		if err != nil {
			return err
		}
		for _, t := range tasks {
			taskOut = append(taskOut, fmt.Sprintf("ID: %d - %s", t.ID, t.Name))
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		// loga o erro raiz de qual serviço quebrou
		log.Printf("GetProjectDetail error: %v", err)
		http.Error(w, "aggregation error: "+err.Error(), http.StatusBadGateway)
		return
	}

	if err := g.Wait(); err != nil {
		http.Error(w, "aggregation error: "+err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Projeto: %s\n", projName)
	fmt.Fprintf(w, "Chaves:\n")
	for _, line := range keysOut {
		fmt.Fprintf(w, "\t%s\n", line)
	}
	fmt.Fprintf(w, "Inventarios:\n")
	for _, line := range invOut {
		fmt.Fprintf(w, "\t%s\n", line)
	}
	fmt.Fprintf(w, "Repositorios:\n")
	for _, line := range repOut {
		fmt.Fprintf(w, "\t%s\n", line)
	}
	fmt.Fprintf(w, "Tarefas:\n")
	for _, line := range taskOut {
		fmt.Fprintf(w, "\t%s\n", line)
	}
}
