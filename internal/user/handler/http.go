package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AndroDeMohawk/notes-api/internal/user/dto"
	"github.com/AndroDeMohawk/notes-api/internal/user/usecase"
	"github.com/AndroDeMohawk/notes-api/pkg/middlewares"
	"github.com/AndroDeMohawk/notes-api/pkg/response"
)

type UserHandler struct {
	userUC *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUC: uc,
	}
}

// POST /auth/register
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input dto.RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.SendJson(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userUC.Register(r.Context(), input)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.SendJson(w, user, http.StatusCreated)
}

// POST /auth/login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.SendJson(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.userUC.Login(r.Context(), input)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response.SendJson(w, map[string]string{"token": token}, http.StatusOK)
}

// GET /users/me
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		response.SendJson(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.userUC.GetProfile(r.Context(), userID)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusNotFound)
		return
	}

	response.SendJson(w, user, http.StatusOK)
}

// PUT /users/me
func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserID(r.Context())
	if !ok {
		response.SendJson(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var input dto.UpdateProfileInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.SendJson(w, "invalid request body", http.StatusBadRequest)
		return
	}

	input.UserID = userID // Подставляем проверенный ID из токена

	user, err := h.userUC.UpdateProfile(r.Context(), input)
	if err != nil {
		response.SendJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.SendJson(w, user, http.StatusOK)
}
