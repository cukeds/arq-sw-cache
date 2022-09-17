package main

import (
	"cache_test/app"
	"cache_test/cache"
	"cache_test/db"
)

func main() {
	cache.Init_cache()
	db.StartDbEngine()
	app.StartRoute()

}
