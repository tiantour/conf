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
	Gateway
	IDC
	Image
	Ipmi
	Mafengwo
	Mchpay
	Modbus
	Nats
	Qiniu
	QQ
	Scene
	S7
	Serial
	Service
	Snmp
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
