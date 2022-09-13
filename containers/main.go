package main

import (
	"cache_test/app"
	"cache_test/db"
	"cache_test/memcached"
)

func main() {
	memcached.Init_cache()
	db.StartDbEngine()
	app.StartRoute()

}
