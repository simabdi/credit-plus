package main

import (
	"credit-plus/internal/config"
	"credit-plus/route"
)

func main() {
	config.Initialize()
	config.Connection()
	route.Initialize()
}
