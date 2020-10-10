package main

import (
	"foreign-currency-go/db"
	"foreign-currency-go/routes"
)

func main() {
	db.Init()
	routes.RunRoutes()
}
