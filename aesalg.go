package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

const KeyLen = 128

// Go\src\crypto\cipher\example_test.go
func EncryptGcm(msg, key string) string {
	if len(key) == 0 {
		return ""
	}
	if len(msg) == 0 {
		return ""
	}
	msgbs := []byte(msg)
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda") // 12 bytes

	keyhash := hashKey(key)
	fmt.Printf("Key Length: %d, Msg length: %d\n", len(keyhash), len(msgbs))
	block, err := aes.NewCipher(keyhash)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, msgbs, nil)
	fmt.Printf("%d: [%x]\n", len(ciphertext), ciphertext)

	res := hex.EncodeToString(ciphertext)
	return res
}

func DecryptGcm(msg, key string) string {
	keyhash := hashKey(key)
	fmt.Printf("Key Length: %d\n", len(keyhash))
	block, err := aes.NewCipher(keyhash)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext, _ := hex.DecodeString(msg)
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda") // 12 bytes
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
	return string(plaintext)
}

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
func EncryptMsg(msg, key string) string {
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
	return hex.EncodeToString(ciphertext)
}

func DecryptMsg(msg, key string) string {
	keyhash := hashKey(key)
	salt := hashKey("abc123!@#")

	block, err := aes.NewCipher(keyhash)
	if err != nil {
		panic(err.Error())
	}

	ciphertext, _ := hex.DecodeString(msg)
	blk := cipher.NewCBCDecrypter(block, salt)
	plain2 := make([]byte, len(ciphertext))
	blk.CryptBlocks(plain2, ciphertext)

	return string(unPaddingPKCS7(plain2))
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

	var cc CharConv
	cc.Load(`SongciRank.txt`)

	s := cc.Encrypt(ciphertext)
	return s
}

func DecryptFromSongci(msg, key string) string {
	var cc CharConv
	cc.Load(`SongciRank.txt`)
	ciphertext := cc.Decrypt(msg)

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

// Keysize: 16|32 bytes, Plaintext: (16|32)*N bytes
func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
	return s
}

func EncryptAESPadding(key []byte, plaintext string) string {
	c, err := aes.NewCipher(key)
	CheckError(err)

	textbs := []byte(plaintext)
	plainPad := paddingPKCS7(textbs, len(key))

	out := make([]byte, len(plainPad))

	c.Encrypt(out, plainPad)
	fmt.Printf("Out: [%X]\n", out)

	return hex.EncodeToString(out)
}

func DecryptAESPadding(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	//fmt.Println("DECRYPTED:", s)
	//unpad := unPaddingPKCS7(pt)
	return string(s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func encdecDemo() {
	keyhash := hashKey(KeyDemo)

	enc := EncryptAESPadding(keyhash, "112233")
	fmt.Println(enc)
	dec := DecryptAESPadding(keyhash, enc)
	fmt.Println(dec)
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
