package repository

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (x *Repository) ReserveProducts(productCodes []string) (bool, error) {
	log.Infof("trying to reserve products with codes=%s", productCodes)

	query := fmt.Sprintf("select * from %s(ARRAY [%s])", reserveProducts, strSliceToPgArray(productCodes))

	var isReservingSuccessful bool

	err := x.db.Get(&isReservingSuccessful, query)
	if err != nil {
		return false, err
	}

	log.Infof("reserving products %s result: %t", productCodes, isReservingSuccessful)
	return isReservingSuccessful, nil
}
