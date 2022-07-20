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
	"limit":         "20",                    // 视频流最多展示的视频数量
	"video_prefix":  "",                      // 视频url前缀，不填则自动设置为本地ip
	"cover_prefix":  "",                      // 封面url前缀，不填则自动设置为本地ip
	"video_dir_fmt": "./public/video/%s%s",   // 视频存放目录。
	"cover_dir_fmt": "./public/cover/%s.png", // 封面存放目录。
}

// Secret 密钥，用于加密的盐
var Secret = "test-secret-credential"

// Port 监听端口
var Port = ":8080"

// NetEnv 网络环境，internal表示内网，external表示外网
var NetEnv = "internal"
