package domain

import "time"

type Payment struct {
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customer_id"`
	Status     string  `json:"status"`
	OrderId    int64   `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  int64   `json:"created_at"`
}

func NewPayment(customerID, OrderId int64, totalPrice float32) Payment {
	return Payment{
		CustomerID: customerID,
		Status:     "Pending",
		OrderId:    OrderId,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now().Unix(),
	}
}
