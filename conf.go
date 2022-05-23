package conf

import (
	"github.com/spf13/viper"
	"github.com/tiantour/conf/iot"
	"github.com/tiantour/conf/storage"
	"github.com/tiantour/conf/trade"
	"github.com/tiantour/conf/union"
	"github.com/tiantour/conf/x"
)

var config Config

type Config struct {
	Token   Token   `toml:"token"`   //
	Gateway Gateway `toml:"gateway"` //
	Service Service `toml:"service"` //

	Cache map[string]storage.Cache `toml:"cache"` //
	DB    map[string]storage.DB    `toml:"db"`    //

	IDC      map[string]iot.IDC      `toml:"idc"`      //
	Ethernet map[string]iot.Ethernet `toml:"ethernet"` //
	Serial   map[string]iot.Serial   `toml:"serial"`   //
	Command  map[string]iot.Command  `toml:"command"`  //
	GPIO     []*iot.GPIO             `toml:"gpio"`     // gpio
	IPMI     []*iot.IPMI             `toml:"ipmi"`     // ipmi
	RTU      []*iot.Modbus           `toml:"rtu"`      // modbus rtu
	S7       []*iot.S7               `toml:"s7"`       // s7
	Snmp     []*iot.Snmp             `toml:"snmp"`     // snmp
	TCP      []*iot.Modbus           `toml:"tcp"`      // modbus tcp

	Alipay map[string]trade.Alipay `toml:"alipay"` //
	Mchpay map[string]trade.Mchpay `toml:"mchpay"` //
	Umspay map[string]trade.Umspay `toml:"umspay"` //
	Wxpay  map[string]trade.Wxpay  `toml:"wxpay"`  //

	QQ     map[string]union.QQ     `toml:"qq"`     //
	Wechat map[string]union.Wechat `toml:"wechat"` //
	Weibo  map[string]union.Weibo  `toml:"weibo"`  //
	Wxwork map[string]union.Wxwork `toml:"wxwork"` //

	Aliyun   map[string]x.Aliyun   `toml:"aliyun"`   //
	Mafengwo map[string]x.Mafengwo `toml:"mafengwo"` //
	Qiniu    map[string]x.Qiniu    `toml:"qiniu"`    //
	Weather  map[string]x.Weather  `toml:"weather"`  //
}

func Init(path, name string, args ...string) error {
	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath(path)
	v.SetConfigName(name)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	for _, item := range args {
		v.SetConfigName(item)
		err = v.MergeInConfig()
		if err != nil {
			continue
		}
	}

	return v.Unmarshal(&config)
}

func NewConfig() *Config {
	return &config
}
