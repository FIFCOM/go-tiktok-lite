package dao

import (
	"crypto/sha256"
	"fmt"
	"github.com/FIFCOM/go-tiktok-lite/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	cfg := config.DB // 导入数据库配置
	conn := cfg["user"] + ":" + cfg["password"] + "@tcp(" + cfg["host"] + ":" + cfg["port"] + ")/" + cfg["dbname"] + "?charset=utf8mb4"
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(conn))
	Handle(err)
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
