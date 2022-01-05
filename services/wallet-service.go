package services

// #cgo CFLAGS: -I/wallet-core/include
// #cgo LDFLAGS: -L/wallet-core/build -L/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWString.h>
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
import "C"
import (
	"errors"
    "github.com/ValeriD/wallet-service/entities"
	"github.com/ValeriD/wallet-service/helpers"
	"encoding/hex"
)

type WalletService interface{


	/**
	* Generating an derrived address for the specified code
	* @param coinType Satoshi's number for the coin
	* @param 
	*/
	GenerateAddress(coinType uint32 , addressIndex uint32 ) entities.Address


	GenerateKeyPair(coinType C.enum_TWCoinType , addressIndex C.uint32_t) (string,string)
}

// Implementation of the interface
type walletService struct {
	wallet *C.struct_TWHDWallet
}

// Constructor
func New(mnemonic string, passPhrase string) (WalletService, error) {
	wallet := CreateWalletWithMnemonic(mnemonic, passPhrase)
	if wallet == nil {
		return nil, errors.New("Invalid mnemonic!")
	}
	return &walletService{
		wallet: wallet,
	}, nil
}

// Implementaions of the methods in the interface
func CreateWalletWithMnemonic(mnemonic string, empty string) *C.struct_TWHDWallet {

	twMnemonic := helpers.ConvertGoStringToTWString(mnemonic)
	twEmpty := helpers.ConvertGoStringToTWString(empty)

	
	return C.TWHDWalletCreateWithMnemonic(twMnemonic, twEmpty)


}

func CreateWallet() (*C.struct_TWHDWallet, error){
	twPassPhrase := helpers.ConvertGoStringToTWString("")

	wallet := C.TWHDWalletCreate(128, twPassPhrase)

	if wallet == nil{
		return nil, errors.New("Unable to create new wallet!")
	}
	return wallet, nil
}


func (service *walletService) GenerateAddress(coinType uint32 , addressIndex uint32 ) entities.Address{
	privateKey, publicKey := service.GenerateKeyPair(coinType, C.uint32_t(addressIndex))

	return entities.Address{
						PublicKey: publicKey,
						PrivateKey: privateKey, 
	}
}

func (service *walletService) GenerateKeyPair(coinType C.enum_TWCoinType , addressIndex C.uint32_t) (string,string){
	privateKey := C.TWHDWalletGetKeyBIP44(service.wallet, coinType, 0 ,0 , addressIndex)

	privateKeyData := C.TWPrivateKeyData(privateKey)

	privateKeyString := hex.EncodeToString(helpers.ConvertTWDataToGoBytes(privateKeyData))

	publicKey := C.TWCoinTypeDeriveAddress(coinType, privateKey)

	publicKeyString := helpers.ConvertTWStringGoToString(publicKey)
	return privateKeyString, publicKeyString
}


