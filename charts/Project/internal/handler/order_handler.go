package handler

import (
	"domain"
)

func (C *domain.CreateOrderInput) CreateOrder(UserID string, ProductID string, Quantity int) (string, error) {
	if C.Quantity <= 0 {
		return "", nil
	}

	return "", nil
}

func ChangeStatus() {

}

func CancelOrder() {

}
