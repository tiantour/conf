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
	Token   map[string]Token   `mapstructure:"token"`   // token
	Gateway map[string]Gateway `mapstructure:"gateway"` // gateway
	Service map[string]Service `mapstructure:"service"` // service

	Cache map[string]storage.Cache `mapstructure:"cache"` // cache
	DB    map[string]storage.DB    `mapstructure:"db"`    // db

	IDC      map[string]iot.IDC      `mapstructure:"idc"`      // idc
	Ethernet map[string]iot.Ethernet `mapstructure:"ethernet"` // ethernet
	Serial   map[string]iot.Serial   `mapstructure:"serial"`   // serial
	Command  map[string]iot.Command  `mapstructure:"command"`  // command
	GPIO     []*iot.GPIO             `mapstructure:"gpio"`     // gpio
	IPMI     []*iot.IPMI             `mapstructure:"ipmi"`     // ipmi
	RTU      []*iot.Modbus           `mapstructure:"rtu"`      // modbus rtu
	S7       []*iot.S7               `mapstructure:"s7"`       // s7
	Snmp     []*iot.Snmp             `mapstructure:"snmp"`     // snmp
	TCP      []*iot.Modbus           `mapstructure:"tcp"`      // modbus tcp

	Alipay map[string]trade.Alipay `mapstructure:"alipay"` // alipay
	Mchpay map[string]trade.Mchpay `mapstructure:"mchpay"` // mchpay
	Umspay map[string]trade.Umspay `mapstructure:"umspay"` // umspay
	Wxpay  map[string]trade.Wxpay  `mapstructure:"wxpay"`  // wxpay

	QQ     map[string]union.QQ     `mapstructure:"qq"`     // qq
	Wechat map[string]union.Wechat `mapstructure:"wechat"` // wechat
	Weibo  map[string]union.Weibo  `mapstructure:"weibo"`  // weibo
	Wxwork map[string]union.Wxwork `mapstructure:"wxwork"` // wxwork

	Aliyun   map[string]x.Aliyun   `mapstructure:"aliyun"`   // aliyun
	Mafengwo map[string]x.Mafengwo `mapstructure:"mafengwo"` // mafengwo
	Qiniu    map[string]x.Qiniu    `mapstructure:"qiniu"`    // qiniu
	Weather  map[string]x.Weather  `mapstructure:"weather"`  // weather
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
