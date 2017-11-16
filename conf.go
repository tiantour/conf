package conf

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	data = &Conf{}
	conf = flag.String("conf", "", "set config file name")
)

// Conf conf
type Conf struct {
	Server   map[string]server
	DB       db
	Cache    cache
	Token    token
	Image    map[string]image
	Wechat   map[string]wechat
	Weibo    weibo
	QQ       qq
	Alipay   alipay
	Wxpay    map[string]wxpay
	Dayu     map[string]dayu
	Weather  weather
	Template map[string]template
	Free     map[string]service
	Manager  map[string]service
	User     map[string]service
}
type server struct {
	IP   string
	Port string
	Host string
	Name string
}
type db struct {
	IP       string
	Port     string
	Database string
	Uname    string
	Passwd   string
}
type cache struct {
	IP   string
	Port string
}
type token struct {
	Secret string
	Issuer string
}
type image struct {
	Bucket string
	Host   string
	Uname  string
	Passwd string
}
type wechat struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}
type weibo struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}
type qq struct {
	AppID  string `toml:"app_id"`
	AppKey string `toml:"app_key"`
}
type alipay struct {
	AppID      string `toml:"app_id"`
	PublicKey  string `toml:"public_key"`
	PrivateKey string `toml:"private_key"`
}
type wxpay struct {
	MchID string `toml:"mch_id"`
	Key   string `toml:"key"`
}
type dayu struct {
	AppKey    string `toml:"app_key"`
	AppSecret string `toml:"app_secret"`
	Sign      string
}
type weather struct {
	ID  string
	Key string
}
type service struct {
	Name string
	Port string
}
type template struct {
	ID   string
	Name string
}

func init() {
	flag.Parse()
	path := "conf.toml"
	if *conf != "" {
		path = fmt.Sprintf("../conf/conf.%s.toml", *conf)
	}
	_, err := toml.DecodeFile(path, data)
	if err != nil {
		fmt.Print(err)
		flag.Usage()
		os.Exit(1)
	}
}

// NewConf new conf
func NewConf() *Conf {
	return data
}
