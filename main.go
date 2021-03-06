package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listTodosRe   = regexp.MustCompile(`^\/todos[\/]*$`)
	getTodosRe    = regexp.MustCompile(`^\/todos\/(\d+)$`)
	createTodosRe = regexp.MustCompile(`^\/todos[\/]*$`)
	deleteTodosRe = regexp.MustCompile(`^\/todos\/(\d+)$`)
)

type todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type datastore struct {
	m map[string]todo
	*sync.RWMutex
}

type objectHandler struct {
	store *datastore
}

func (h *objectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listTodosRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getTodosRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createTodosRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodDelete && deleteTodosRe.MatchString(r.URL.Path):
		h.Delete(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *objectHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	todos := make([]todo, 0, len(h.store.m))
	for _, v := range h.store.m {
		todos = append(todos, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *objectHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getTodosRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	u, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Todo not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// Create a new todo
func (h *objectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u todo
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// Delete todo with id
func (h *objectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	matches := deleteTodosRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.Lock()
	delete(h.store.m, matches[1])
	h.store.Unlock()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo deleted"))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
	mux := http.NewServeMux()
	todoH := &objectHandler{
		store: &datastore{
			m: map[string]todo{
				"1": {ID: "1", Title: "watch go programming tutorials"},
				"2": {ID: "2", Title: "learn ansible and docker"},
				"3": {ID: "3", Title: "read go programming books"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/todos", todoH)
	mux.Handle("/todos/", todoH)

	// Set up a server listening on port 8080 with the mux as the handler
	// If you set on localhost:8080, you can access the server from your browser locally
	// but your container will not be able to make it accessible from outside
	// http.ListenAndServe("0.0.0.0:8080", mux)
	http.ListenAndServe(":8080", mux)
}
