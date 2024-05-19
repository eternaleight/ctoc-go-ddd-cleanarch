package models

import "time"

type Purchase struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"id"`
	Price           int       `gorm:"column:price" json:"price"`
	ProductID       uint      `gorm:"column:productId;index" json:"productId"`
	BuyerID         uint      `gorm:"column:buyerId;index" json:"buyerId"`
	PurchaseDate    time.Time `gorm:"column:purchaseDate" json:"purchaseDate"`
	StripePaymentID string    `gorm:"column:stripePaymentId" json:"stripePaymentId"`
}

func (Purchase) TableName() string {
	return "Purchase"
}
