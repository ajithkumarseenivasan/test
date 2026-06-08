package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management/model"
	"user-management/service"
)

type VendorHandler struct {
	service service.VendorService
}

func NewVendorHandler(service service.VendorService) *VendorHandler {
	return &VendorHandler{service: service}
}

func (h *VendorHandler) SaveVendor(w http.ResponseWriter, r *http.Request) {
	var req model.VendorRequest

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

	resp, err := h.service.SaveVendor(req)
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

func (h *VendorHandler) GetVendors(w http.ResponseWriter, r *http.Request) {
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
	description := query.Get("description")
	email := query.Get("email")
	contactPersonName := query.Get("contactPersonName")

	vendors, total, err := h.service.GetVendors(tenantId, name, description, email, contactPersonName, page, limit)
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
			Content: model.VendorListResponse{
				Vendors: vendors,
				Page:    page,
				Limit:   limit,
				Total:   total,
			},
			Message: model.Success,
		},
	)
}
