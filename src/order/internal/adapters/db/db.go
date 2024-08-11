package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpc-microservices-go/order/internal/application/core/domain"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItems
}

type OrderItems struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderId     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&Order{}, &OrderItems{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db}, nil
}

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	result := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}
	return order, result.Error
}

func (a Adapter) Save(order domain.Order) error {
	var orderItems []OrderItems
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItems{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}
