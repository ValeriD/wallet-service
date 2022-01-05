package controllers

import (
	"github.com/gin-gonic/gin"
	 . "github.com/ValeriD/wallet-service/services"
	//  "github.com/ValeriD/wallet-service/entities"
	"strconv"
	"net/http"
	"errors"
)

var (
	errInvalidAddressIndex  = errors.New("Invalid address index")
	errInvalidBody     		= errors.New("Invalid request body")
)

type WalletController interface {
	generateAddress(ctx *gin.Context, coinType uint32)
	GenerateBitcoinAddress(ctx *gin.Context) 
	GenerateEthereumAddress(ctx *gin.Context) 
}

type walletController struct {
	service WalletService
}

func New(service WalletService) WalletController {

	return &walletController{
		service: service,
	}
}

func (walletController *walletController) generateAddress(ctx *gin.Context, coinType uint32){
	addressIndexString := ctx.Query("addressIndex")

	addressIndexUint64, err := strconv.ParseUint(addressIndexString, 10, 32)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": "failed", "message": errInvalidAddressIndex.Error()})	
		return
	}

	address := walletController.service.GenerateAddress(coinType, uint32(addressIndexUint64))

	ctx.JSON(http.StatusOK, gin.H{ "status": "success", "address":address}) 
}

func (walletController *walletController) GenerateBitcoinAddress(ctx *gin.Context) {
	walletController.generateAddress(ctx, 0);
}

func (walletController *walletController) GenerateBitcoinTestnetAddress(ctx *gin.Context) {
	walletController.generateAddress(ctx, 1);
}

func (walletController *walletController) GenerateEthereumAddress(ctx *gin.Context) {
	walletController.generateAddress(ctx, 60);
}
//To add other coins just create a new method, that calls generateAddress