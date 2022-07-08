package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/FIFCOM/go-tiktok-lite/config"
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"github.com/jxskiss/base62"
	"io"
	"log"
	"strconv"
)

// GetToken 生成用户token
func GetToken(user dao.User) string {
	token, err := Encrypt([]byte(strconv.FormatInt(user.Id, 10)), []byte(Hash(config.Secret)))
	Handle(err)
	return token
}

// ParseToken 解析用户token，返回dao.User
func ParseToken(token string) (user dao.User, err error) {
	id, err := Decrypt(token, []byte(Hash(config.Secret)))
	Handle(err)
	user.Id, err = strconv.ParseInt(string(id), 10, 64)
	Handle(err)
	return dao.GetUserById(user.Id)
}

func Hash(s string) string {
	// return sha256(s + config.Secret)
	s += config.Secret
	hash := sha256.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Encrypt encrypts data using 256-bit AES-GCM, return base62 encoded string
func Encrypt(plain []byte, key []byte) (ciphered string, err error) {
	k := sha256.Sum256(key)
	block, err := aes.NewCipher(k[:])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	sealed := gcm.Seal(nonce, nonce, plain, nil)
	ciphered = base62.EncodeToString(sealed) // base62 encode sealed text
	return ciphered, nil
}

// Decrypt decode data using base62 and decrypt data using 256-bit AES-GCM, return base62 encoded string
func Decrypt(ciphered string, key []byte) (plain string, err error) {
	k := sha256.Sum256(key)                          // sha256 hash key
	ciphertext, err := base62.DecodeString(ciphered) // base62 decode ciphered text
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(k[:])
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}
	opened, err := gcm.Open(nil, ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():], nil)
	return string(opened), err
}

func Handle(e error) {
	if e != nil {
		log.Panicf("[ERR] Tiktok DAO Layer Error : %v", e)
	}
}
