package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const path = "public/conf/conf.toml"

// Options Options
var Options options

type (
	// options
	options struct {
		Server  server  `toml:"server"`
		DB      db      `toml:"db"`
		Account account `toml:"account"`
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
		RAS     rsa     `toml:"rsa"`
	}
	// server
	server struct {
		Host   string
		Port   string
		Upload string
		Debug  bool
	}
	// db
	db struct {
		Uname    string
		Passwd   string
		Database string
	}
	// account
	account struct {
		Secret string
		Issuer string
		Avatar string
		Gender string
		Salt   string
	}
	// upaiyun
	upyun struct {
		Bucket   string
		Username string
		Passwd   string
		Host     string
	}
	// qiniu
	qiniu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		Host      string
	}
	// weather
	weather struct {
		ID  string
		Key string
	}
	// alidayu
	alidayu struct {
		AppKey      string
		AppSecret   string
		SmsTemplate string
		SmsSign     string
	}
	// wechat
	wechat struct {
		AppID     string
		AppSecret string
		Token     string
	}
	// weibo
	weibo struct {
		AppID     string
		AppSecret string
	}
	// qq
	qq struct {
		AppID     string
		AppSecret string
	}
	// wxpay
	wxpay struct {
		AppID     string
		AccountID string
		MchID     string
		NotifyURL string
		Key       string
	}
	// alipay
	alipay struct {
		AppID     string
		Partner   string
		Key       string
		Seller    string
		NotifyURL string
	}
	// rsa
	rsa struct {
		PublicKey  string
		PrivateKey string
	}
	// unique
	unique struct {
		Table string
	}
)

// 初始化
func init() {
	_, err := toml.DecodeFile(path, &Options)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
