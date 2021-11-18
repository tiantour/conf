package conf

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	conf data
)

type data struct {
	Alipay
	Aliyun
	Cache
	CB
	Command
	DB
	Ethernet
	Gateway
	GPIO
	IDC
	Mafengwo
	Mchpay
	Point
	Qiniu
	QQ
	Scene
	Serial
	Service
	Token
	Umspay
	Weather
	Wechat
	Weibo
	Wxpay
	Wxwork
}

func init() {
	path := flag.String("path", "conf.toml", "config file path")
	flag.Parse()

	_, err := toml.DecodeFile(*path, &conf)
	if err != nil {
		log.Println(err)
		flag.Usage()
		os.Exit(1)
	}
}
