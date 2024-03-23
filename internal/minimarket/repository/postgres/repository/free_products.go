package repository

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (x *Repository) FreeProducts(productCodes []string) (bool, error) {
	log.Infof("trying to free products with codes=%s", productCodes)

	query := fmt.Sprintf("select * from %s(ARRAY [%s])", freeProducts, strSliceToPgArray(productCodes))

	var isSuccessful bool

	err := x.db.Get(&isSuccessful, query)
	if err != nil {
		return false, err
	}

	log.Infof("free products %s result: %t", productCodes, isSuccessful)

	return isSuccessful, nil
}
