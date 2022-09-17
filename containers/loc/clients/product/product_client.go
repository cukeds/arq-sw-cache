package product

import (
	"cache_test/model"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type productClient struct{}

type ProductClientInterface interface {
	GetProductById(id int) model.Product
	UpdateProduct(id int, desc string) model.Product
}

var (
	ProductClient ProductClientInterface
)

func init() {
	ProductClient = &productClient{}
}

func (s *productClient) GetProductById(id int) model.Product {
	var product model.Product
	Db.Where("product_id = ?", id).First(&product)
	return product
}

func (s *productClient) UpdateProduct(id int, desc string) model.Product {
	var product model.Product

	Db.Model(&product).Where("product_id = ?", id).Update("description", desc)

	return product

}
