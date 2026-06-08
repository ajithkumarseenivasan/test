package service

import (
	"time"
	"user-management/model"
	"user-management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VendorService interface {
	SaveVendor(vendor model.VendorRequest) (bool, error)
	GetVendors(tenantId string, name string, description string, email string, contactPersonName string, page int64, limit int64) ([]model.Vendor, int64, error)
}

type vendorService struct {
	repo repository.VendorRepository
}

func NewVendorService(r repository.VendorRepository) VendorService {
	return &vendorService{repo: r}
}

func (s *vendorService) SaveVendor(vendor model.VendorRequest) (bool, error) {
	vendor.Vendor.TenantID, _ = primitive.ObjectIDFromHex(vendor.TenantId)
	vendor.Vendor.CreatedBy, _ = primitive.ObjectIDFromHex(vendor.UserId)
	vendor.Vendor.ModifiedDate = time.Now().UTC()
	vendor.Vendor.CreatedDate = time.Now().UTC()
	return s.repo.SaveVendor(vendor.Vendor)
}

func (s *vendorService) GetVendors(tenantId string, name string, description string, email string, contactPersonName string, page int64, limit int64) ([]model.Vendor, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit
	return s.repo.GetVendors(tenantId, name, description, email, contactPersonName, skip, limit)
}
