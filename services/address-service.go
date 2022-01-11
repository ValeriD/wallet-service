package services

// #cgo CFLAGS: -I/wallet-core/include
// #cgo LDFLAGS: -L/wallet-core/build -L/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWCoinTypeConfiguration.h>
import "C"
import (
	"errors"
    "github.com/ValeriD/wallet-service/entities"
	"github.com/ValeriD/wallet-service/helpers"
)

type AddressService interface{
	/**
	* Generating an derrived address for the specified code
	* @param coinType Satoshi's number for the coin
	* @param addressIndex - used by for derivation of the wallet. If it is used once it will return the previous generated address, else will create  a new derived address
	* @return new Address object
	*/
	GenerateAddress(coinType uint32 , addressIndex uint32 ) entities.Address

	/**
	* Generating a pair of public and private keys for the provided coin type and address index
	* @param coinType 
	* @param addressIndex - used by for derivation of the wallet. If it is used once it will return the previous generated address, else will create  a new derived address
	* @return public and private keys 
	*/
	generatePublicKey(coinType C.enum_TWCoinType , addressIndex C.uint32_t) string
}

// Implementation of the interface
type AddressService struct {
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
//////////////////////////////////////////////////
// Implementaions of the methods in the interface
//////////////////////////////////////////////////

func (service *AddressService) GenerateAddress(coinType uint32 , addressIndex uint32 ) entities.Address{
	publicKey := service.generatePublicKey(coinType, C.uint32_t(addressIndex))

	return entities.Address{
						PublicKey: publicKey,
						CoinType: helpers.ConvertTWStringToGoString(C.TWCoinTypeConfigurationGetID(coinType)), 
	}
}

func (service *AddressService) generatePublicKey(coinType C.enum_TWCoinType , addressIndex C.uint32_t) string{
	privateKey := C.TWHDWalletGetKeyBIP44(service.wallet, coinType, 0 ,0 , addressIndex)

	publicKey := C.TWCoinTypeDeriveAddress(coinType, privateKey)

	publicKeyString := helpers.ConvertTWStringToGoString(publicKey)
	return publicKeyString
}


