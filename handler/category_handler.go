package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	resp, err := h.service.SaveCategory(req)
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

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.ParseInt(query.Get("page"), 10, 64)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.ParseInt(query.Get("limit"), 10, 64)
	if err != nil || limit < 1 {
		limit = 10
	}
	tenantId := query.Get("tenantId")

	categories, total, err := h.service.GetCategories(tenantId, page, limit)
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
			Status: true,
			Content: model.CategoryListResponse{
				Categories: categories,
				Page:       page,
				Limit:      limit,
				Total:      total,
			},
			Message: model.Success,
		},
	)
}
