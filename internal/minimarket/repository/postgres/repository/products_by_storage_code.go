package repository

import (
	"fmt"

	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (x *Repository) GetProductsByStorageCode(storageCode string) ([]model.Product, error) {
	log.Infof("trying to get products by storage code=%s", storageCode)

	query := fmt.Sprintf("select * from %s('%s')", selectProductsByStorageCode, storageCode)

	var products []model.Product
	err := x.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	log.Info("products: ", products)
	return products, nil
}
