package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"

	"github.com/bluewd111/go-template/app/user/command"
	"github.com/bluewd111/go-template/app/user/query"
	"github.com/bluewd111/go-template/app/user/service"
)

const USER_PATH = "/users/"

var instance *UserHandler
var once sync.Once

type UserHandler struct {
	userService *service.UserService
}

func GetUserHandlerInstance(userService *service.UserService) *UserHandler {
	// 複数インスタンス生成、ルーティングをさせたくないので Singleton にしている
	once.Do(func() {
		instance = &UserHandler{
			userService: userService,
		}
	})
	return instance
}

func (h *UserHandler) GetPathAndHandler() (string, http.HandlerFunc) {
	return USER_PATH, h.handlers()
}

func (h *UserHandler) handlers() http.HandlerFunc {
	pathParam := regexp.MustCompile(`^/users/(.+)+?/`)

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			matches := pathParam.FindStringSubmatch(r.URL.Path)
			if matches == nil {
				h.getUsersHandler(w, r)
				return
			}
			h.getUserHandler(w, r, matches[0])
		case http.MethodPost:
			h.createUserHandler(w, r)
		case http.MethodPatch:
			h.updateUserHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (h *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var cmd command.CreateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var cmd command.UpdateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateUser(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.userService.GetUsers())
}

func (h *UserHandler) getUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	var query query.GetUserQuery
	query.ID = id
	if query.ID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(query)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
