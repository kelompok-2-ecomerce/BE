package handler

import (
	"projects/features/cart"
)

type AddUpdateCartRequest struct {
	ItemID uint   `form:"itemid"`
	Qty    string `form:"qty"`
}

func (data *AddUpdateCartRequest) reqToCore() cart.Core {
	return cart.Core{
		Qty:    data.Qty,
		ItemID: data.ItemID,
	}
}

func ToCore(data interface{}) *cart.Core {
	res := cart.Core{}

	switch data.(type) {
	case AddUpdateCartRequest:
		cnv := data.(AddUpdateCartRequest)
		res.Qty = cnv.Qty
		res.ItemID = cnv.ItemID

	default:
		return nil
	}

	return &res
}
