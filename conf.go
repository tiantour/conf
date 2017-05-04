package conf

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type (
	data struct {
		Server   server   `toml:"server"`
		DB       db       `toml:"db"`
		Account  account  `toml:"account"`
		Upyun    upyun    `toml:"upyun"`
		Qiniu    qiniu    `toml:"qiniu"`
		Weather  weather  `toml:"weather"`
		Wechat   wechat   `toml:"wechat"`
		Weibo    weibo    `toml:"weibo"`
		QQ       qq       `toml:"qq"`
		Wxpay    wxpay    `toml:"wxpay"`
		Alipay   alipay   `toml:"alipay"`
		RSA      rsa      `toml:"rsa"`
		Alidayu  alidayu  `toml:"alidayu"`
		Template template `toml:"template"`
	}
	// server
	server struct {
		Host   string
		Port   string
		Upload string
		Start  string
	}
	// DB db
	db struct {
		Uname    string
		Passwd   string
		Database string
	}
	// account
	account struct {
		Default int
		Secret  string
		Issuer  string
		Avatar  string
		Gender  string
		Salt    string
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
	// wechat
	wechat struct {
		AppID        string
		AppSecret    string
		WebAppID     string
		WebAppSecret string
		Token        string
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
	// alidayu
	alidayu struct {
		AppKey    string
		AppSecret string
		Template  string
		Sign      string
	}
	rsa struct {
		PublicKey  string
		PrivateKey string
	}
	template struct {
		Feedback string
		Stock    string
		Notice   string
		Warn     string
		Material string
		Report   string
	}
)

// Config config
var (
	Data = new(data)
	conf = flag.String("conf", "", "set config file name")
)

func init() {
	flag.Parse()
	if *conf == "" {
		log.Println("conf file name is null")
		flag.Usage()
		os.Exit(1)
	}
	env := os.Getenv("ENV")
	if env != "" {
		env = fmt.Sprintf(".%s", env)
	}
	path := fmt.Sprintf(
		"../conf/conf.%s%s.toml",
		*conf,
		env,
	)
	_, err := toml.DecodeFile(path, Data)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
