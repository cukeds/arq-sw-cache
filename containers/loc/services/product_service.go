package services

import (
	pclient "cache_test/clients/product"
	"cache_test/dto"
	"cache_test/model"
	e "cache_test/utils/errors"
)

type productService struct {
	productClient pclient.ProductClientInterface
}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	UpdateProduct(id int, desc string) (dto.ProductDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func initProductService(productClient pclient.ProductClientInterface) productServiceInterface {
	service := new(productService)
	service.productClient = productClient
	return service
}

func init() {
	ProductService = initProductService(pclient.ProductClient)
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product

	product = s.productClient.GetProductById(id)
	var productDto dto.ProductDto

	if product.ProductId < 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}

	productDto.ProductId = product.ProductId
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Price = product.Price
	productDto.CurrencyId = product.CurrencyId
	productDto.Stock = product.Stock
	productDto.Picture = product.Picture

	return productDto, nil
}

func (s *productService) UpdateProduct(id int, desc string) (dto.ProductDto, e.ApiError) {

	var product model.Product = s.productClient.UpdateProduct(id, desc)

	var productDto dto.ProductDto
	productDto.ProductId = product.ProductId
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Price = product.Price
	productDto.CurrencyId = product.CurrencyId
	productDto.Stock = product.Stock
	productDto.Picture = product.Picture

	return productDto, nil
}
