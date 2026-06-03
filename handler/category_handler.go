package handler

import (
	"encoding/json"
	"net/http"
	"user-management/model"
	"user-management/service"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) SaveCategory(w http.ResponseWriter, r *http.Request) {
	var req model.CategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(
			model.MasterUiResponse{
				Status:  false,
				Content: err.Error(),
				Message: model.Failed,
			},
		)
		return
	}

	resp, err := h.service.SaveCategory(req.Category)
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
			Content: resp,
			Message: model.Success,
		},
	)
}
