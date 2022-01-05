package main

import (
	"github.com/ValeriD/wallet-service/services"
	"github.com/ValeriD/wallet-service/controllers"
	"log"
	"github.com/gin-gonic/gin"
)

func main(){
	log.Println("Starting server...")

	server := gin.Default()

	walletService, err := services.New("confirm bleak useless tail chalk destroy horn step bulb genuine attract split", "")
	if err != nil || walletService == nil {
		log.Fatal(err)
		return
	}
	
	walletController := controllers.New(walletService)


	// Routes
	initRoutes(server, walletController)


	server.Run(":8080")
}

func initRoutes(router *gin.Engine, controller controllers.WalletController) {
	router.GET("/generate-bitcoin-address", controller.GenerateBitcoinAddress)
	router.GET("/generate-ethereum-address", controller.GenerateEthereumAddress)
}
