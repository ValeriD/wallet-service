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

func  ConvertTWStringToGoString(s unsafe.Pointer) string {
    return C.GoString(C.TWStringUTF8Bytes(s))
}

func ConvertGoBytesToTWData(d []byte) unsafe.Pointer {
    cBytes := C.CBytes(d)
    defer C.free(unsafe.Pointer(cBytes))
    data := C.TWDataCreateWithBytes((*C.uchar)(cBytes), C.ulong(len(d)))
    return data
}

func  ConvertTWDataToGoBytes(d unsafe.Pointer) []byte {
    cBytes := C.TWDataBytes(d)
    cSize := C.TWDataSize(d)
    return C.GoBytes(unsafe.Pointer(cBytes), C.int(cSize))
}