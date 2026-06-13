package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management/model"
	"user-management/service"
)

type InventoryItemHandler struct {
	service service.InventoryItemService
}

func NewInventoryItemHandler(service service.InventoryItemService) *InventoryItemHandler {
	return &InventoryItemHandler{service: service}
}

func (h *InventoryItemHandler) SaveInventoryItem(w http.ResponseWriter, r *http.Request) {
	var req model.InventoryItemRequest

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

	resp, err := h.service.SaveInventoryItem(req)
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

func (h *InventoryItemHandler) GetInventoryItems(w http.ResponseWriter, r *http.Request) {
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
	name := query.Get("name")
	vendorId := query.Get("vendorId")

	inventoryItems, total, err := h.service.GetInventoryItems(tenantId, name, vendorId, page, limit)
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
			Content: model.InventoryItemListResponse{
				InventoryItems: inventoryItems,
				Page:           page,
				Limit:          limit,
				Total:          total,
			},
			Message: model.Success,
		},
	)
}
