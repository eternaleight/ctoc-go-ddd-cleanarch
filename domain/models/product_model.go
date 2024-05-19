package models

import "time"

type Product struct {
	ID          uint       `gorm:"primaryKey;column:id" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Price       int        `gorm:"column:price" json:"price"`
	ImageURL    string     `gorm:"column:imageUrl" json:"imageUrl"`
	ImageSize   uint64     `gorm:"column:imageSize" json:"imageSize"`
	VideoURL    string     `gorm:"column:videoUrl" json:"videoUrl"`
	CreatedAt   time.Time  `gorm:"column:createdAt" json:"createdAt"`
	SellerID    uint       `gorm:"column:sellerId;index" json:"sellerId"`
	Purchases   []Purchase `gorm:"foreignKey:ProductID" json:"purchases"`
}

func (Product) TableName() string {
	return "Product"
}
