package config

// DB 数据库配置
var DB = map[string]string{
	"host":     "localhost", // 数据库地址
	"port":     "3306",      // 数据库端口
	"user":     "root",      // 数据库用户名
	"password": "123456",    // 数据库密码
	"dbname":   "tiktok",    // 数据库名
}

// Video 视频配置
var Video = map[string]string{
	"limit":        "20",                            // 视频流最多展示的视频数量
	"video_prefix": "https://api.fifcom.cn/vid/",    // 视频url前缀
	"cover_prefix": "https://www.fifcom.cn/avatar/", // 封面url前缀
}

// Secret 密钥，用于加密的盐
var Secret = "test-secret-credential"
