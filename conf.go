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
	Cache
	CB
	Dayu
	DB
	Device
	Image
	Mafengwo
	Mchpay
	Point
	Power
	QQ
	Sensor
	Server
	Service
	Token
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
