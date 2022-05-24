package iot

// Snmp snmp
type Snmp struct {
	Model     int32  `mapstructure:"model"`     // 型号
	Device    string `mapstructure:"device"`    // 设备
	Port      string `mapstructure:"port"`      // 端口
	IP        string `mapstructure:"ip"`        // 地址
	Community string `mapstructure:"community"` // 社区
	Uint      string `mapstructure:"unit"`      // 单位
}
