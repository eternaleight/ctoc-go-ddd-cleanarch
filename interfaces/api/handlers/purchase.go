package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/domain/models"
	"github.com/gin-gonic/gin"
)

type PurchaseUsecasesInterface interface {
	CreatePurchase(purchase *models.Purchase) error
	GetPurchaseByID(id uint) (*models.Purchase, error)
}

// 購入関連のハンドラを管理
type PurchaseHandler struct {
	PurchaseUsecases PurchaseUsecasesInterface
}

// 新しいPurchaseHandlerを初期化して返す
func NewPurchaseHandler(purchaseUsecases PurchaseUsecasesInterface) *PurchaseHandler {
	return &PurchaseHandler{
		PurchaseUsecases: purchaseUsecases,
	}
}

// 新しい購入を作成するためのハンドラ
func (ph *PurchaseHandler) CreatePurchase(c *gin.Context) {
	var purchase models.Purchase

	// 購入データのJSONをパース
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "購入データの形式が正しくない"})
		return
	}

	// 購入データをデータベースに保存
	err := ph.PurchaseUsecases.CreatePurchase(&purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースに購入情報を保存できなかった"})
		return
	}

	// 保存に成功した場合のレスポンスを返す
	c.JSON(http.StatusOK, gin.H{"data": "商品の購入が成功した"})
}
