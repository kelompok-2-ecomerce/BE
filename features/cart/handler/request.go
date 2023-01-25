package handler

import "projects/features/cart"

type AddCartReq struct {
	Qty int `json:"qty" form:"qty"`
}

func ToCore(data interface{}) *cart.Core {
	res := cart.Core{}

	switch data.(type) {
	case AddCartReq:
		cnv := data.(AddCartReq)
		res.Qty = cnv.Qty
	default:
		return nil
	}

	return &res
}
