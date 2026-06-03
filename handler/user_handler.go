package handler

import (
	"encoding/json"
	"net/http"
	"user-management/model"
	"user-management/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		json.NewEncoder(w).Encode(
			model.MasterUiResponse{
				Status:  false,
				Content: err.Error(),
				Message: model.Failed,
			},
		)
		return
	}
	json.NewEncoder(w).Encode(
		model.MasterUiResponse{
			Status:  true,
			Content: users,
			Message: model.Success,
		},
	)
}

func (h *UserHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	userName := requestParams["name"]

	user, err := h.service.GetUserByName(userName)
	if err != nil {
		json.NewEncoder(w).Encode(
			model.MasterUiResponse{
				Status:  false,
				Content: err.Error(),
				Message: model.Failed,
			},
		)
	} else {
		json.NewEncoder(w).Encode(
			model.MasterUiResponse{
				Status:  true,
				Content: user,
				Message: model.Success,
			},
		)
	}
}

// func (h *UserHandler) SaveNewUser(w http.ResponseWriter, r *http.Request) {
// 	var req model.MasterUiRequest

// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(
// 			model.MasterUiResponse{
// 				Status:  false,
// 				Content: err.Error(),
// 				Message: model.Failed,
// 			},
// 		)
// 	}

// 	resp, err := h.service.SaveUser(req.User)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(
// 			model.MasterUiResponse{
// 				Status:  false,
// 				Content: err.Error(),
// 				Message: model.Failed,
// 			},
// 		)
// 	}

// 	json.NewEncoder(w).Encode(
// 		model.MasterUiResponse{
// 			Status:  true,
// 			Content: resp,
// 			Message: model.Success,
// 		},
// 	)
// }
