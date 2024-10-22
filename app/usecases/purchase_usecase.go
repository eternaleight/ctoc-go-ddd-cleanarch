package usecases

import (
	"github.com/eternaleight/go-backend/domain/models"
)

// 購入ストア操作のインターフェースを定義
type PurchaseStoreInterface interface {
	CreatePurchase(purchase *models.Purchase) error
	GetPurchaseByID(id uint) (*models.Purchase, error)
}

type PurchaseUsecases struct {
	PurchaseStore PurchaseStoreInterface
}

// PurchaseUsecasesの新しいインスタンスを初期化
func NewPurchaseUsecases(purchaseStore PurchaseStoreInterface) *PurchaseUsecases {
	return &PurchaseUsecases{
		PurchaseStore: purchaseStore,
	}
}

func (u *PurchaseUsecases) CreatePurchase(purchase *models.Purchase) error {
	// 購入データをデータベースに保存
	err := u.PurchaseStore.CreatePurchase(purchase)
	if err != nil {
		return err
	}

	return nil
}

func (u *PurchaseUsecases) GetPurchaseByID(id uint) (*models.Purchase, error) {
	// 指定されたIDの購入情報を取得
	purchase, err := u.PurchaseStore.GetPurchaseByID(id)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}
