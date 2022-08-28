package gnscrypto

// #cgo CFLAGS: -I../../openssl/include -I../../gnscryptocpp -I/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX12.1.sdk/usr/include/usr/inc
// #cgo LDFLAGS: -L../../darwin_build/gnscrypto/lib -lgnscrypto -lz -lc++
// #include "../../gnscryptocpp/gns_crypto.h"
import "C"
import (
	"unsafe"
)

func clen(n []byte) int {
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			return i
		}
	}
	return len(n)
}
func Cchar2string(data unsafe.Pointer, len int) string {
	//Convert raw C++ unchar string to golang string
	bytes := C.GoBytes(data, (C.int)(len))
	full := string(bytes)
	return full
}

func Cchar2byte(data unsafe.Pointer, len C.ulong) []byte {
	bytes := C.GoBytes(data, (C.int)(len))
	return bytes
}

func GeneratePubKeyFromPriv(cert_name string) string {

	name := C.CString(cert_name)
	x := C.malloc(C.sizeof_uchar * 64)
	defer C.free(unsafe.Pointer(x))
	y := C.malloc(C.sizeof_uchar * 64)
	defer C.free(unsafe.Pointer(y))
	C.generate_pubkey_from_privkey(name, (*C.uchar)(x), (*C.uchar)(y), (C.bool)(false))

	x_bytes := C.GoBytes(x, 64)
	x_str := string(x_bytes)
	y_bytes := C.GoBytes(y, 64)
	y_str := string(y_bytes)
	//fmt.Println(x_str)
	//fmt.Println(y_str)
	return x_str + y_str
}

func GenerateChallenge() ([]byte, string) {
	//std::string c_string = string_to_char_array(input);
	challenge_len := C.ulong(256)
	digest_len := C.ulong(64)
	raw_challenge := C.malloc(C.sizeof_uchar * challenge_len)
	defer C.free(unsafe.Pointer(raw_challenge))
	digest := C.malloc(C.sizeof_uchar * digest_len)
	defer C.free(unsafe.Pointer(digest))

	C.generate_challenge((*C.uchar)(raw_challenge), (*C.uchar)(digest), (C.bool)(false))

	//raw_challenge_str := Cchar2byte(raw_challenge, challenge_len)
	//digest_str := Cchar2byte(raw_challenge, digest_len)

	raw_challenge_bytes := C.GoBytes(raw_challenge, 256)
	digest_bytes := C.GoBytes(digest, 64)

	digest_string := string(digest_bytes)

	return raw_challenge_bytes, digest_string
}

func SignChallenge(challenge []byte, cert_name string) string {

	name := C.CString(cert_name)
	r := C.malloc(C.sizeof_uchar * 64)
	defer C.free(unsafe.Pointer(r))
	s := C.malloc(C.sizeof_uchar * 64)
	defer C.free(unsafe.Pointer(s))

	C.sign_challenge((*C.uchar)(unsafe.Pointer(&challenge[0])), C.ulong(len(challenge)), name, (*C.char)(r), (*C.char)(s), (C.bool)(false))

	return Cchar2string(r, 64) + Cchar2string(s, 64)
	//return "hello"
}

func VerifySignature(s_challenge []byte, challenge []byte, cert []byte) int {
	//(s_challenge) should be 64 bytes
	//(challenge) should be 256 bytes
	//(cert) should be 377 bytes

	verified := C.int(0)
	//fmt.Println("about to call C verify")

	C.ECDSA_dss_verify_signature((*C.int)(unsafe.Pointer(&verified)), (*C.uchar)(unsafe.Pointer(&s_challenge[0])), (*C.uchar)(unsafe.Pointer(&challenge[0])),
		(*C.uchar)(unsafe.Pointer(&cert[0])), C.int(len(cert)))

	return int(verified)
}

func EncodeData(input string, ephemeral []byte, cert_name string) string {

	if len(input) < 1 {
		return ""
	}
	name := C.CString(cert_name)
	encrypt_key := C.malloc(C.sizeof_uchar * 16)
	encrypt_iv := C.malloc(C.sizeof_uchar * 16)
	encrypted_output := C.malloc(C.sizeof_uchar * 2048)
	c_input := C.CString(input)
	output_len := C.int(2048)

	//fmt.Println("Token:")
	//fmt.Println(ephemeral)
	C.ECDH((*C.uchar)(unsafe.Pointer(&ephemeral[0])), name, (*C.uchar)(unsafe.Pointer(encrypt_key)), (*C.uchar)(unsafe.Pointer(encrypt_iv)))
	//fmt.Println(Cchar2string(encrypt_key, 16))
	C.AES128_Encrypt((*C.uchar)(unsafe.Pointer(c_input)), (C.int)(len(input)), (*C.uchar)(unsafe.Pointer(encrypt_key)), (*C.uchar)(unsafe.Pointer(encrypt_iv)),
		(*C.uchar)(unsafe.Pointer(encrypted_output)), &output_len, (C.bool)(false))
	//fmt.Println("Encryted length:")
	//fmt.Println(output_len)

	defer C.free(unsafe.Pointer(encrypt_key))
	defer C.free(unsafe.Pointer(encrypt_iv))
	defer C.free(unsafe.Pointer(encrypted_output))
	//defer C.free(unsafe.Pointer(c_input))

	return Cchar2string(encrypted_output, (int)(output_len))
}

func DecodeData(input []byte, ephemeral []byte, cert_name string) string {
	if len(input) < 1 {
		return ""
	}
	name := C.CString(cert_name)
	encrypt_key := C.malloc(C.sizeof_uchar * 16)
	encrypt_iv := C.malloc(C.sizeof_uchar * 16)
	decrypted_output := C.malloc(C.sizeof_uchar * 2048)
	output_len := C.int(2048)
	C.ECDH((*C.uchar)(unsafe.Pointer(&ephemeral[0])), name, (*C.uchar)(encrypt_key), (*C.uchar)(encrypt_iv))
	C.AES128_Decrypt((*C.uchar)(unsafe.Pointer(&input[0])), (C.int)(len(input)),
		(*C.uchar)(encrypt_key), (*C.uchar)(encrypt_iv), (*C.uchar)(decrypted_output), &output_len, (C.bool)(false))
	defer C.free(unsafe.Pointer(encrypt_key))
	defer C.free(unsafe.Pointer(encrypt_iv))
	defer C.free(unsafe.Pointer(decrypted_output))
	return Cchar2string(decrypted_output, (int)(output_len))
}
