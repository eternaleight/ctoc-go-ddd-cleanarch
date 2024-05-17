package usecases

import (
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
)

type PurchaseUsecasesInterface interface {
	CreatePurchase(purchase *models.Purchase) error
	GetPurchaseByID(id uint) (*models.Purchase, error)
}

type PurchaseUsecases struct {
	PurchaseStore stores.PurchaseStoreInterface
}

// NewPurchaseUsecasesはPurchaseUsecasesの新しいインスタンスを初期化します
func NewPurchaseUsecases(purchaseStore stores.PurchaseStoreInterface) *PurchaseUsecases {
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
