package helpers

// #cgo CFLAGS: -I/wallet-core/include
// #cgo LDFLAGS: -L/wallet-core/build -L/wallet-core/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include<TrustWalletCore/TWString.h>
// #include<TrustWalletCore/TWData.h>
import "C"
import "unsafe"



func ConvertGoStringToTWString(str string) unsafe.Pointer{
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	resStr := C.TWStringCreateWithUTF8Bytes(cStr)
	return resStr
}

func  ConveretTWDataToGoBytes(data unsafe.Pointer) []byte{
	cBytes := C.TWDataBytes(data)
    cSize := C.TWDataSize(data)
    return C.GoBytes(unsafe.Pointer(cBytes), C.int(cSize))
}

func  ConvertTWDataToGoBytes(d unsafe.Pointer) []byte {
    cBytes := C.TWDataBytes(d)
    cSize := C.TWDataSize(d)
    return C.GoBytes(unsafe.Pointer(cBytes), C.int(cSize))
}

func  ConvertTWStringGoToString(s unsafe.Pointer) string {
    return C.GoString(C.TWStringUTF8Bytes(s))
}