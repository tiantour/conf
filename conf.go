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
	Token   map[string]Token   `toml:"token"`   // token
	Gateway map[string]Gateway `toml:"gateway"` // gateway
	Service map[string]Service `toml:"service"` // service

	Cache map[string]storage.Cache `toml:"cache"` // cache
	DB    map[string]storage.DB    `toml:"db"`    // db

	IDC      map[string]iot.IDC      `toml:"idc"`      // idc
	Ethernet map[string]iot.Ethernet `toml:"ethernet"` // ethernet
	Serial   map[string]iot.Serial   `toml:"serial"`   // serial
	Command  map[string]iot.Command  `toml:"command"`  // command
	GPIO     []*iot.GPIO             `toml:"gpio"`     // gpio
	IPMI     []*iot.IPMI             `toml:"ipmi"`     // ipmi
	RTU      []*iot.Modbus           `toml:"rtu"`      // modbus rtu
	S7       []*iot.S7               `toml:"s7"`       // s7
	Snmp     []*iot.Snmp             `toml:"snmp"`     // snmp
	TCP      []*iot.Modbus           `toml:"tcp"`      // modbus tcp

	Alipay map[string]trade.Alipay `toml:"alipay"` // alipay
	Mchpay map[string]trade.Mchpay `toml:"mchpay"` // mchpay
	Umspay map[string]trade.Umspay `toml:"umspay"` // umspay
	Wxpay  map[string]trade.Wxpay  `toml:"wxpay"`  // wxpay

	QQ     map[string]union.QQ     `toml:"qq"`     // qq
	Wechat map[string]union.Wechat `toml:"wechat"` // wechat
	Weibo  map[string]union.Weibo  `toml:"weibo"`  // weibo
	Wxwork map[string]union.Wxwork `toml:"wxwork"` // wxwork

	Aliyun   map[string]x.Aliyun   `toml:"aliyun"`   // aliyun
	Mafengwo map[string]x.Mafengwo `toml:"mafengwo"` // mafengwo
	Qiniu    map[string]x.Qiniu    `toml:"qiniu"`    // qiniu
	Weather  map[string]x.Weather  `toml:"weather"`  // weather
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
