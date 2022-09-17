package product

import (
	"cache_test/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
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
	log.Debug("Product: ", product)

	return product
}

func (s *productClient) UpdateProduct(id int, desc string) model.Product {
	var product model.Product

	Db.Model(&product).Where("product_id = ?", id).Update("description", desc)
	
	log.Debug("Product: ", product)
	return product

}
