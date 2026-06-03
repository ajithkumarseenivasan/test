package handler

import (
	"encoding/json"
	"net/http"
	"time"
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
	// Validate MasterUiRequest fields (provided at top-level via embedding)
	if req.TenantId == "" || req.UserId == "" {
		json.NewEncoder(w).Encode(
			model.MasterUiResponse{
				Status:  false,
				Content: "TenantId and UserId are required",
				Message: model.Failed,
			},
		)
		return
	}

	// Merge master fields into category, but don't overwrite provided values.
	// if req.Category.TenantID == "" {
	// 	req.Category.TenantID = req.TenantId
	// }
	if req.Category.CreatedBy == "" {
		req.Category.CreatedBy = req.UserId
	}
	if req.Category.CreatedDate.IsZero() {
		req.Category.CreatedDate = time.Now()
	}
	if req.Category.ModifiedDate.IsZero() {
		req.Category.ModifiedDate = time.Now()
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
