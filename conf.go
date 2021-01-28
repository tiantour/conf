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
	Future
	Gateway
	Gear
	GPIO
	IDC
	Image
	Liquid
	Mafengwo
	Mchpay
	Next
	Point
	Power
	Qiniu
	QQ
	Scene
	Sensor
	Serial
	Server
	Service
	Target
	Token
	Umspay
	Weather
	Wechat
	Weibo
	Wxpay
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
