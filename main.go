package main

import (
	"github.com/lucassouzz/gopportunities/config"
	"github.com/lucassouzz/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	router.Initialize()
}
