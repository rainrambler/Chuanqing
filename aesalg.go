package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"golang.org/x/crypto/sha3"
)

const KeyLen = 128

// 16 bytes
func hashKey(key string) []byte {
	if len(key) == 0 {
		return []byte{}
	}
	bskey := []byte(key)
	digest := sha3.Sum256(bskey)
	return digest[:KeyLen/8]
}

// https://asecuritysite.com/golang/go_symmetric
func EncryptToSongci(msg, key string) string {
	keyhash := hashKey(key)
	salt := hashKey("abc123!@#")

	block, err := aes.NewCipher(keyhash)
	if err != nil {
		panic(err.Error())
	}

	textbs := []byte(msg)
	plainPad := paddingPKCS7(textbs, len(keyhash))
	ciphertext := make([]byte, len(plainPad))

	blk := cipher.NewCBCEncrypter(block, salt)
	blk.CryptBlocks(ciphertext, plainPad)

	s := convInst.Encrypt(ciphertext)
	return s
}

func DecryptFromSongci(msg, key string) string {
	ciphertext := convInst.Decrypt(msg)

	keyhash := hashKey(key)
	salt := hashKey("abc123!@#")

	block, err := aes.NewCipher(keyhash)
	if err != nil {
		panic(err.Error())
	}

	//ciphertext, _ := hex.DecodeString(msg)
	blk := cipher.NewCBCDecrypter(block, salt)
	plain2 := make([]byte, len(ciphertext))
	blk.CryptBlocks(plain2, ciphertext)

	return string(unPaddingPKCS7(plain2))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// https://juejin.cn/post/7190571062025781285
// PKCS7 填充
func paddingPKCS7(plaintext []byte, blockSize int) []byte {
	paddingSize := blockSize - len(plaintext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	bs := append(plaintext, paddingText...)
	//fmt.Printf("Orig: [%X], Pad: [%X], block: %d\n", plaintext, bs, blockSize)
	return bs
}

// PKCS7 反填充
func unPaddingPKCS7(s []byte) []byte {
	length := len(s)
	if length == 0 {
		return s
	}
	unPadding := int(s[length-1])
	return s[:(length - unPadding)]
}
