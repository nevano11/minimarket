package model

import (
	"fmt"
	"strings"
)

type OrderType int

const (
	None OrderType = iota
	Asc
	Desc
)

type PostFilter struct {
	PageNumber *int      `json:"page_number"`
	PageSize   *int      `json:"page_size"`
	PriceOrder OrderType `json:"price_order"`
	DateOrder  OrderType `json:"date_order"`
	MaxPrice   *int      `json:"max_price"`
	MinPrice   *int      `json:"min_price"`
}

func (f PostFilter) String() string {
	sb := strings.Builder{}

	sb.WriteString("PostFilter:[")

	if f.PriceOrder != None {
		sb.WriteString(fmt.Sprintf("PriceOrder: %d,", f.PriceOrder))
	}
	if f.DateOrder != None {
		sb.WriteString(fmt.Sprintf("DateOrder: %d,", f.DateOrder))
	}
	if f.MaxPrice != nil {
		sb.WriteString(fmt.Sprintf("MaxPrice: %d,", *f.MaxPrice))
	}
	if f.MinPrice != nil {
		sb.WriteString(fmt.Sprintf("MinPrice: %d,", *f.MinPrice))
	}
	if f.PageNumber != nil {
		sb.WriteString(fmt.Sprintf("PageNumber: %d,", *f.PageNumber))
	}
	if f.PageSize != nil {
		sb.WriteString(fmt.Sprintf("PageSize: %d,", *f.PageSize))
	}

	sb.WriteString("]")

	return sb.String()
}
