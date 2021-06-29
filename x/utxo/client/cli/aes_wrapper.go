package cli

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func Encode(key string, data string) string {

	//16,24,32 key length
	var aeskey = []byte(md5_32(key))

	dataByte := []byte(data)

	xpass, _ := AesEncrypt(dataByte, aeskey)

	pass64 := base64.StdEncoding.EncodeToString(xpass)

	return pass64
}

//对字符串进行MD5哈希，返回 32 位结果
func md5_32(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func main() {

	var aeskey = []byte("1234567887654321")

	pass64 := Encode("1234567887654321", "mydata")

	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)

}
