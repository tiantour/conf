package conf

import (
	"log"

	"github.com/spf13/viper"
	"github.com/tiantour/conf/v2/iot"
	"github.com/tiantour/conf/v2/iot/port"
	"github.com/tiantour/conf/v2/iot/protocol"
	"github.com/tiantour/conf/v2/storage"
	"github.com/tiantour/conf/v2/trade"
	"github.com/tiantour/conf/v2/union"
	"github.com/tiantour/conf/v2/x"
)

var (
	// Path path
	Path string

	// Data data
	Data *Config
)

type Config struct {
	Token map[string]Token         // token
	Cache map[string]storage.Cache // cache
	DB    map[string]storage.DB    // db

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
	Gateway  map[string]x.Gateway  // gateway
	Service  map[string]x.Service  // service

	IDC      map[string]iot.IDC       // idc
	Scene    map[string]iot.Scene     // scene
	Command  map[string]iot.Command   // command
	Ethernet map[string]port.Ethernet // ethernet
	Serial   map[string]port.Serial   // serial
	GPIO     []*port.GPIO             // gpio
	IPMI     []*protocol.IPMI         // ipmi
	RTU      []*protocol.Modbus       // modbus rtu
	S7       []*protocol.S7           // s7
	Snmp     []*protocol.Snmp         // snmp
	TCP      []*protocol.Modbus       // modbus tcp
}

func New(args ...string) {
	v := viper.New()
	v.SetConfigName("common")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	v.AddConfigPath(Path)

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config err: %v", err)
	}

	for _, item := range args {
		v.SetConfigName(item)
		err := v.MergeInConfig()
		if err != nil {
			log.Fatalf("merge config err: %v", err)
		}
	}

	err = v.Unmarshal(&Data)
	if err != nil {
		log.Fatalf("parse config err: %v", err)
	}
}
