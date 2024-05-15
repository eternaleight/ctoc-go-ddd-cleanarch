// handlers/products.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductStore stores.ProductStoreInterface
}

// 新しいProductHandlerのインスタンスを初期化
func NewProductHandler(productStore stores.ProductStoreInterface) *ProductHandler {
	return &ProductHandler{
		ProductStore: productStore,
	}
}

// 新しい商品を作成
func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var product usecases.ProductInput
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品のデータの形式が正しくありません。"})
		return
	}

	createdProduct, err := usecases.CreateProduct(ph.ProductStore, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースに商品を保存できませんでした。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": createdProduct})
}

// 全商品をリストとして取得
func (ph *ProductHandler) ListProducts(c *gin.Context) {
	products, err := usecases.ListProducts(ph.ProductStore)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品のリストの取得に失敗しました。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// 指定されたIDの商品を取得
func (ph *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	product, err := usecases.GetProductByID(ph.ProductStore, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の情報を取得できませんでした。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// 指定されたIDの商品情報を更新
func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	var product usecases.ProductInput
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品のデータの形式が正しくありません。"})
		return
	}

	updatedProduct, err := usecases.UpdateProduct(ph.ProductStore, uint(id), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の更新に失敗しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedProduct})
}

// 指定されたIDの商品を削除
func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	err = usecases.DeleteProduct(ph.ProductStore, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の削除に失敗しました。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "商品は正常に削除されました。"})
}
