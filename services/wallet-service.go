package services
// #cgo CFLAGS: -I/wallet-core/include
// #cgo LDFLAGS: -L/wallet-core/build -L/wallet
// #include <TrustWalletCore/TWHDWallet.h>
import "C"
import(
	"github.com/ValeriD/wallet-service/helpers"
)

type WalletService interface{
	CreateWallet()
	CreateWalletWithMnemonic()
}

type WalletService struct{}

func New() *WalletService {
	return &WalletService{}
}

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