package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Options 设置
var Options config

// Config 配置文件
type config struct {
	Server  server  `toml:"server"`
	DB      db      `toml:"db"`
	Upyun   upyun   `toml:"upyun"`
	Qiniu   qiniu   `toml:"qiniu"`
	Weather weather `toml:"weather"`
	Alidayu alidayu `toml:"alidayu"`
	Wechat  wechat  `toml:"wechat"`
	Weibo   weibo   `toml:"weibo"`
	QQ      qq      `toml:"qq"`
	Wxpay   wxpay   `toml:"wxpay"`
	Alipay  alipay  `toml:"alipay"`
	Unique  unique  `toml:"unique"`
}

// 服务器
type server struct {
	Host   string
	Port   string
	Upload string
}

// 数据库
type db struct {
	Uname    string
	Passwd   string
	Database string
}

// upaiyun
type upyun struct {
	Bucket   string
	Username string
	Passwd   string
	Host     string
}

// qiniu
type qiniu struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Host      string
}

// weather
type weather struct {
	ID  string
	Key string
}

// alidayu
type alidayu struct {
	AppKey      string
	AppSecret   string
	SmsTemplate string
	SmsSign     string
}

// wechat
type wechat struct {
	AppID     string
	AppSecret string
	Token     string
}

// weibo
type weibo struct {
	AppID     string
	AppSecret string
}

// qq
type qq struct {
	AppID     string
	AppSecret string
}

// wxpay
type wxpay struct {
	AppID     string
	AccountID string
	MchID     string
	NotifyURL string
	Key       string
}

// alipay
type alipay struct {
	AppID     string
	Partner   string
	Key       string
	Seller    string
	NotifyURL string
}

// unique
type unique struct {
	Table string
}

// 初始化
func init() {
	_, err := toml.DecodeFile("public/conf/conf.toml", &Options)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
