package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management/model"
	"user-management/service"
)

type PurchaseOrderHandler struct {
	service service.PurchaseOrderService
}

func NewPurchaseOrderHandler(service service.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{service: service}
}

func (h *PurchaseOrderHandler) SavePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	var req model.PurchaseOrderRequest

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

	resp, err := h.service.SavePurchaseOrder(req)
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

func (h *PurchaseOrderHandler) GetPurchaseOrders(w http.ResponseWriter, r *http.Request) {
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
	orderId := query.Get("orderId")
	vendorName := query.Get("vendorName")
	status := query.Get("status")

	purchaseOrders, total, err := h.service.GetPurchaseOrders(tenantId, status, vendorName, orderId, page, limit)
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
			Content: model.PurchaseOrderListResponse{
				PurchaseOrders: purchaseOrders,
				Page:           page,
				Limit:          limit,
				Total:          total,
			},
			Message: model.Success,
		},
	)
}
