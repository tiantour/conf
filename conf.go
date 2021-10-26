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
	Gpio
	IDC
	Ipmi
	Mafengwo
	Mchpay
	ModbusRTU
	ModbusTCP
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
