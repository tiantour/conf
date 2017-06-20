package conf

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	conf    = flag.String("conf", "", "set config file name")
	newConf = new(Conf)
)

type (
	// Conf conf
	Conf struct {
		Server   server   `toml:"server"`
		DB       db       `toml:"db"`
		Cache    cache    `toml:"cache"`
		Account  account  `toml:"account"`
		Upyun    upyun    `toml:"upyun"`
		Qiniu    qiniu    `toml:"qiniu"`
		Weather  weather  `toml:"weather"`
		Wechat   wechat   `toml:"wechat"`
		Weibo    weibo    `toml:"weibo"`
		QQ       qq       `toml:"qq"`
		Wxpay    wxpay    `toml:"wxpay"`
		Alipay   alipay   `toml:"alipay"`
		Alidayu  alidayu  `toml:"alidayu"`
		Template template `toml:"template"`
	}
	// server
	server struct {
		Domain string
		Port   string
		Upload string
		Start  string
	}
	// DB db
	db struct {
		Uname    string
		Passwd   string
		Host     string
		Port     string
		Database string
	}
	// cache
	cache struct {
		Host string
		Port string
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
	mp struct {
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
		AppID      string
		Method     string
		Gateway    string
		PublicKey  string
		PrivateKey string
		NotifyURL  string
	}
	// alidayu
	alidayu struct {
		AppKey    string
		AppSecret string
		Template  string
		Sign      string
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

// NewConf new conf
func NewConf() *Conf {
	return newConf
}
func init() {
	flag.Parse()
	if *conf == "" {
		log.Println("conf file is null")
		flag.Usage()
		os.Exit(1)
	}
	env := os.Getenv("ENV")
	if env != "" {
		env = fmt.Sprintf("%s.", env)
	}
	path := fmt.Sprintf("../conf/conf.%s.%stoml", *conf, env)
	_, err := toml.DecodeFile(path, newConf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
