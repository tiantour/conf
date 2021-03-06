package conf

type Snmp struct {
	Data []*SnmpItem `toml:"snmp"`
}

// NewSnmp new snmp
func NewSnmp() *Snmp {
	return &conf.Snmp
}

// Snmp snmp item
type SnmpItem struct {
	Model     int32  `toml:"model"`     // 型号
	Device    string `toml:"device"`    // 设备
	Port      string `toml:"port"`      // 端口
	IP        string `toml:"ip"`        // 地址
	Community string `toml:"community"` // 社区
	Uint      string `toml:"unit"`      // 单位
}
