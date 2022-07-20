package dao

import (
	"encoding/json"
	"github.com/FIFCOM/go-tiktok-lite/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var DB *gorm.DB

func init() {
	var err error
	cfg := config.DB // 导入数据库配置
	// 设置编码为utf8以及设置解析时间格式（如果不设置解析时间则会将时间转换为字符串导致报错）
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	conn := cfg["user"] + ":" + cfg["password"] + "@tcp(" + cfg["host"] + ":" + cfg["port"] + ")/" + cfg["dbname"] + "?charset=utf8mb4&parseTime=true"
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(conn))
	Handle(err)

	// 填写Video配置，通过修改config.NetEnv来设置是使用内网IP还是外网IP
	if config.Video["video_prefix"] == "" {
		if config.NetEnv == "internal" {
			config.Video["video_prefix"] = "http://" + getInternalIP() + config.Port + "/douyin/video/"
		} else {
			config.Video["video_prefix"] = "http://" + getExternalIP() + config.Port + "/douyin/video/"
		}
	}
	if config.Video["cover_prefix"] == "" {
		if config.NetEnv == "internal" {
			config.Video["cover_prefix"] = "http://" + getInternalIP() + config.Port + "/douyin/cover/"
		} else {
			config.Video["cover_prefix"] = "http://" + getExternalIP() + config.Port + "/douyin/cover/"
		}
	}
}

// Handle 处理错误
func Handle(e error) {
	if e != nil {
		log.Panicf("[ERR] Tiktok DAO Layer Error : %v", e)
	}
}

// getInternalIP 获取内网IP
func getInternalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	Handle(err)
	defer func(conn net.Conn) {
		err = conn.Close()
		Handle(err)
	}(conn)
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	addr, _, err := net.SplitHostPort(localAddr)
	return addr
}

// getExternalIP 使用API获取外网IP
func getExternalIP() string {
	type IP struct {
		Query string
	}
	req, err := http.Get("http://ip-api.com/json/")
	Handle(err)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		Handle(err)
	}(req.Body)
	body, err := ioutil.ReadAll(req.Body)
	Handle(err)
	var ip IP
	err = json.Unmarshal(body, &ip)
	Handle(err)
	return ip.Query
}
