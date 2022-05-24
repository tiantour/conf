package conf

import (
	"log"

	"github.com/spf13/viper"
	"github.com/tiantour/conf/v2/iot"
	"github.com/tiantour/conf/v2/storage"
	"github.com/tiantour/conf/v2/trade"
	"github.com/tiantour/conf/v2/union"
	"github.com/tiantour/conf/v2/x"
)

const (
	_TYPE = "toml"
	_NAME = "common"
)

var (
	config Config
)

type Config struct {
	Token   map[string]Token   // token
	Gateway map[string]Gateway // gateway
	Service map[string]Service // service

	Cache map[string]storage.Cache // cache
	DB    map[string]storage.DB    // db

	IDC      map[string]iot.IDC      // idc
	Ethernet map[string]iot.Ethernet // ethernet
	Serial   map[string]iot.Serial   // serial
	Command  map[string]iot.Command  // command
	GPIO     []*iot.GPIO             // gpio
	IPMI     []*iot.IPMI             // ipmi
	RTU      []*iot.Modbus           // modbus rtu
	S7       []*iot.S7               // s7
	Snmp     []*iot.Snmp             // snmp
	TCP      []*iot.Modbus           // modbus tcp

	Alipay map[string]trade.Alipay // alipay
	Mchpay map[string]trade.Mchpay // mchpay
	Umspay map[string]trade.Umspay // umspay
	Wxpay  map[string]trade.Wxpay  // wxpay

	QQ     map[string]union.QQ     // qq
	Wechat map[string]union.Wechat // wechat
	Weibo  map[string]union.Weibo  // weibo
	Wxwork map[string]union.Wxwork // wxwork

	Aliyun   map[string]x.Aliyun   // aliyun
	Mafengwo map[string]x.Mafengwo // mafengwo
	Qiniu    map[string]x.Qiniu    // qiniu
	Weather  map[string]x.Weather  // weather
}

func New(path string, name ...string) {
	v := viper.New()
	v.SetConfigType(_TYPE)
	v.SetConfigName(_NAME)
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config err: %v", err)
	}

	for _, item := range name {
		v.SetConfigName(item)
		err := v.MergeInConfig()
		if err != nil {
			log.Fatalf("merge config err: %v", err)
		}
	}

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("parse config err: %v", err)
	}
}

func NewConfig() *Config {
	return &config
}
