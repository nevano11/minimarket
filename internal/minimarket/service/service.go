package service

import (
	"github.com/nevano11/minimarket/internal/minimarket/entity"
)

type Storege interface {
	GetProductsByStorageCode(string) ([]entity.Product, error)
	ReserveProducts([]string) (bool, error)
	FreeProducts([]string) (bool, error)
}

type Service struct {
	storage Storege
}

func NewService(storage Storege) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetProductsByStorageCode(storageCode string) ([]entity.Product, error) {
	return s.storage.GetProductsByStorageCode(storageCode)
}

func (s *Service) ReserveProducts(productCodes []string) (bool, error) {
	return s.storage.ReserveProducts(productCodes)
}

func (s *Service) FreeProducts(productCodes []string) (bool, error) {
	return s.storage.FreeProducts(productCodes)
}
