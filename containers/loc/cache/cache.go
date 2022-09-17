package cache

import (
	"cache_test/dto"
	e "cache_test/utils/errors"
	"fmt"
	ccache "github.com/karlseguin/ccache/v2"
	"strconv"
	"time"
)

var cache *ccache.Cache
var DefaultTTL time.Duration

func Init_cache() {
	cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(500))
	DefaultTTL = 2 * time.Minute
}

func Set(product dto.ProductDto) (dto.ProductDto, e.ApiError) {

	id := strconv.Itoa(product.ProductId)
	cache.Set(id, product, DefaultTTL)
	return product, nil
}

func Get(id string) (value dto.ProductDto, apiError e.ApiError) {
	item := cache.Get(id)
	if item == nil {
		return dto.ProductDto{}, e.NewNotFoundApiError(fmt.Sprintf("product %s not found", id))
	}
	if item.Expired() {
		return dto.ProductDto{}, e.NewNotFoundApiError(fmt.Sprintf("product %s not found", id))
	}
	return item.Value().(dto.ProductDto), nil

}
