package iot

// Command command
type Command struct {
	SlaveID  byte   `mapstructure:"slaveid"`  // 站号 modbus
	Slot     int    `mapstructure:"slot"`     // 插槽 s7
	Function uint16 `mapstructure:"function"` // 功能 modbus&s7
	Address  uint16 `mapstructure:"address"`  // 地址 modbus&s7
	Location uint16 `mapstructure:"location"` // 地址 modbus&s7
	Quantity uint16 `mapstructure:"quantity"` // 数量 modbus&s7
	Value    []byte `mapstructure:"value"`    // 数值 modbus&s7
	Uname    string `mapstructure:"uname"`    // 用户 impi
	Passwd   string `mapstructure:"passwd"`   // 密码 impi
	Code     string `mapstructure:"code"`     // 命令 impi
}
