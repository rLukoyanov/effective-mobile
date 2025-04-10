package main

import (
	"effective-mobile/config"
	"effective-mobile/internal/logger"
)

func main() {
	// init cfg
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// init logger
	logger.InitLogger(cfg.LOG_LEVEL, cfg.MODE)
	// inig storage
	// init echo
	// grasful shotdown
}
