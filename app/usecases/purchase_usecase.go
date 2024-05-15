package usecases

import (
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
)

func CreatePurchase(store stores.PurchaseStoreInterface, purchase *models.Purchase) error {
	// 購入データをデータベースに保存
	err := store.CreatePurchase(purchase)
	if err != nil {
		return err
	}

	return nil
}

func GetPurchaseByID(store stores.PurchaseStoreInterface, id uint) (*models.Purchase, error) {
	// 指定されたIDの購入情報を取得
	purchase, err := store.GetPurchaseByID(id)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}
