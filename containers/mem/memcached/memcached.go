package memcached

import (
	"cache_test/dto"
	e "cache_test/utils/errors"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

var cache *memcache.Client

func Init_cache() {
	cache = memcache.New("cache:11211")
}

func Set(key string, value []byte) {

	cache.Set(&memcache.Item{Key: key, Value: value})
}

func Get(key string) (value dto.ProductDto, apiError e.ApiError) {
	item, err := cache.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return dto.ProductDto{}, e.NewNotFoundApiError(fmt.Sprintf("product %s not found", key))
		}
		return dto.ProductDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting product %s", key), err)
	}

	var pDto dto.ProductDto
	if err := json.Unmarshal(item.Value, &pDto); err != nil {
		return dto.ProductDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting product %s", key), err)
	}

	return pDto, nil
}

func Flush() (apiError e.ApiError) {
	if err := cache.FlushAll(); err != nil {
		return e.NewInternalServerApiError(err.Error(), err)
	}
	return nil

}
