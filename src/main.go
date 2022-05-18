package main

import (
	"Shopify-Challenge/src/database"
	"Shopify-Challenge/src/logger"
	"Shopify-Challenge/src/router"
	"net/http"
)

func main() {

	logger.NewLogger()

	if err := database.DBInit(); err != nil {
		logger.Sugar.Error("failed to launch database: ", err)
		return
	}
	logger.Sugar.Info("database connection established")

	r := router.NewRouter()
	logger.Sugar.Info("api service up and running")
	if err := http.ListenAndServe(":3333", r); err != nil {
		logger.Sugar.Error("failed to launch api server: ", err)
		return
	}

}
