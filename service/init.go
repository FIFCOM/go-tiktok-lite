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
	token, err := Encrypt(strconv.FormatInt(user.Id, 10), Hash(config.Secret))
	Handle(err)
	return token
}

// ParseToken 解析用户token，返回dao.User
func ParseToken(token string) (user dao.User, err error) {
	id, err := Decrypt(token, Hash(config.Secret))
	if id == "" {
		id = "0"
	}
	user.Id, err = strconv.ParseInt(id, 10, 64)
	Handle(err)
	return dao.GetUserById(user.Id)
}

// Hash 哈希字符串
func Hash(s string) string {
	// return sha256(s + config.Secret)
	s += config.Secret
	hash := sha256.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Encrypt 使用AES-GCM-256加密数据，并返回base62编码的字符串
func Encrypt(data string, key string) (ciphered string, err error) {
	plain := []byte(data)
	k := sha256.Sum256([]byte(key))
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
	ciphered = base62.EncodeToString(sealed) // 使用base62编码
	return ciphered, nil
}

// Decrypt 使用AES-GCM-256解密字符串，并返回解密后的原始数据
func Decrypt(ciphered string, key string) (plain string, err error) {
	k := sha256.Sum256([]byte(key))                  // 使用sha256哈希key
	ciphertext, err := base62.DecodeString(ciphered) // 使用base62解码
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
		return "", errors.New("无法解密：无效密文")
	}
	opened, err := gcm.Open(nil, ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():], nil)
	return string(opened), err
}

// Handle 处理错误
func Handle(e error) {
	if e != nil {
		log.Panicf("[ERR] Tiktok Service Layer Error : %v", e)
	}
}
