package handler

import (
	"encoding/json"
	"net/http"
	"user-management/middleware"
	"user-management/model"
	"user-management/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MasterUiResponse{Status: false, Content: err.Error(), Message: model.Failed})
		return
	}

	resp, err := h.service.Register(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MasterUiResponse{Status: false, Content: err.Error(), Message: model.Failed})
		return
	}

	json.NewEncoder(w).Encode(model.MasterUiResponse{Status: true, Content: resp, Message: model.Success})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MasterUiResponse{Status: false, Content: err.Error(), Message: model.Failed})
		return
	}

	resp, err := h.service.Login(req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.MasterUiResponse{Status: false, Content: err.Error(), Message: model.Failed})
		return
	}

	json.NewEncoder(w).Encode(model.MasterUiResponse{Status: true, Content: resp, Message: model.Success})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetAuthClaims(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.MasterUiResponse{Status: false, Content: "unauthorized", Message: model.Failed})
		return
	}

	json.NewEncoder(w).Encode(model.MasterUiResponse{Status: true, Content: claims, Message: model.Success})
}
