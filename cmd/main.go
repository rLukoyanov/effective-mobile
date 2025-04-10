package main

import (
	"context"
	"effective-mobile/config"
	"effective-mobile/internal/db"
	"effective-mobile/internal/logger"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	// init cfg
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// init logger
	logger.InitLogger(cfg.LOG_LEVEL, cfg.MODE)
	// inig storage
	db := db.ConnectDB(cfg, ctx)
	defer db.Close()
	// init echo
	e := echo.New()
	e.Logger.Fatal(e.Start(":" + cfg.PORT))
	// grasful shotdown
}
