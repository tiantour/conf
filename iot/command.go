package iot

// Command command
type Command struct {
	SlaveID  byte   `toml:"slaveid"`  // 站号 modbus
	Slot     int    `toml:"slot"`     // 插槽 s7
	Function uint16 `toml:"function"` // 功能 modbus&s7
	Address  uint16 `toml:"address"`  // 地址 modbus&s7
	Location uint16 `toml:"location"` // 地址 modbus&s7
	Quantity uint16 `toml:"quantity"` // 数量 modbus&s7
	Value    []byte `toml:"value"`    // 数值 modbus&s7
	Uname    string `toml:"uname"`    // 用户 impi
	Passwd   string `toml:"passwd"`   // 密码 impi
	Code     string `toml:"code"`     // 命令 impi
}
