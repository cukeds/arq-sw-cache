package services

import (
	pclient "cache_test/clients/product"
	"cache_test/dto"
	"cache_test/model"
	e "cache_test/utils/errors"

	log "github.com/sirupsen/logrus"
)

type productService struct {
	productClient pclient.ProductClientInterface
}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetProductsByCategoryId(id int) (dto.ProductsDto, e.ApiError)
	GetProductsBySearch(query string) (dto.ProductsDto, e.ApiError)
	GetNProducts(n int) (dto.ProductsDto, e.ApiError)
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

func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = s.productClient.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.ProductId = product.ProductId
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}

func (s *productService) GetNProducts(n int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = s.productClient.GetNProducts(n)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto

		productDto.ProductId = product.ProductId
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}

func (s *productService) GetProductsByCategoryId(id int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = s.productClient.GetProductsByCategoryId(id)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.ProductId = product.ProductId
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}

func (s *productService) GetProductsBySearch(query string) (dto.ProductsDto, e.ApiError) {
	var products model.Products
	products = s.productClient.GetProductsBySearch(query)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.ProductId = product.ProductId
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}
