package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/harunsasmaz/kubernetes-playground/internal/storage"
)

type Handler struct {
	storage *storage.Storage
}

func New(ctx context.Context) *Handler {
	return &Handler{
		storage: storage.New(ctx),
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := h.storage.DB().Ping(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := h.storage.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) GetDone(w http.ResponseWriter, r *http.Request) {
	todos, err := h.storage.GetDone()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) GetRemaining(w http.ResponseWriter, r *http.Request) {
	todos, err := h.storage.GetRemaining()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	todo, err := h.storage.Get(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var todo storage.TODO
	json.NewDecoder(r.Body).Decode(&todo)

	todo.Status = "waiting"
	todo.CreatedAt = time.Now().UTC()

	res, err := h.storage.Create(todo.ID, todo.Title, todo.Description, todo.Status, todo.CreatedAt, todo.DueDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var todo storage.TODO
	json.NewDecoder(r.Body).Decode(&todo)

	res, err := h.storage.UpdateStatus(todo.ID, todo.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := h.storage.Delete(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
