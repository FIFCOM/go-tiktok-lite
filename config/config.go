package config

// DB 数据库配置
var DB = map[string]string{
	"host":     "localhost",
	"port":     "3306",
	"user":     "root",
	"password": "123456",
	"dbname":   "tiktok",
}

// Secret 密钥，用于加密的盐
var Secret = "test-secret-credential"
