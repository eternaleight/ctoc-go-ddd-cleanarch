package stores

import (
	"github.com/eternaleight/go-backend/domain/models"
	"gorm.io/gorm"
)

// PurchaseStoreInterfaceは購入ストア操作のインターフェースを定義します
type PurchaseStoreInterface interface {
	CreatePurchase(purchase *models.Purchase) error
	GetPurchaseByID(id uint) (*models.Purchase, error)
}

// PurchaseStoreは購入に関連するデータベース操作を管理します
type PurchaseStore struct {
	db *gorm.DB
}

// NewPurchaseStoreはPurchaseStoreの新しいインスタンスを初期化します
func NewPurchaseStore(db *gorm.DB) *PurchaseStore {
	return &PurchaseStore{db: db}
}

// CreatePurchaseは新しい購入をデータベースに保存します
func (ps *PurchaseStore) CreatePurchase(purchase *models.Purchase) error {
	return ps.db.Create(purchase).Error
}

// GetPurchaseByIDは指定されたIDの購入情報を取得します
func (ps *PurchaseStore) GetPurchaseByID(id uint) (*models.Purchase, error) {
	var purchase models.Purchase
	err := ps.db.First(&purchase, id).Error
	return &purchase, err
}
