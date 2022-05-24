package iot

// Command command
type Command struct {
	SlaveID  byte   // 站号 modbus
	Slot     int    // 插槽 s7
	Function uint16 // 功能 modbus&s7
	Address  uint16 // 地址 modbus&s7
	Location uint16 // 地址 modbus&s7
	Quantity uint16 // 数量 modbus&s7
	Value    []byte // 数值 modbus&s7
	Uname    string // 用户 impi
	Passwd   string // 密码 impi
	Code     string // 命令 impi
}
