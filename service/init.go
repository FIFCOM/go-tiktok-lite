package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/FIFCOM/go-tiktok-lite/config"
	"github.com/FIFCOM/go-tiktok-lite/dao"
	"log"
)

func GetToken(user dao.User) string {
	return Hash(user.Name + user.Password + config.Secret)
}

func Hash(s string) string {
	// return sha256(s + config.Secret)
	s += config.Secret
	hash := sha256.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Handle(e error) {
	if e != nil {
		log.Panicf("[ERR] Tiktok DAO Layer Error : %v", e)
	}
}
