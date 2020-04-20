package conf

// Command command
type Command struct {
	Data map[string]command `toml:"command"`
}

// NewCommand new command
func NewCommand() *Command {
	return &conf.Command
}

type command struct {
	SlaveID  byte   `toml:"slaveid"`  // 站号
	Location uint16 `toml:"location"` // 地址
	Function uint16 `toml:"function"` // 功能
	Address  uint16 `toml:"address"`  // 地址
	Quantity uint16 `toml:"quantity"` // 数量
	Value    []byte `toml:"value"`    // 数值
	Uname    string `toml:"uname"`    // 用户
	Passwd   string `toml:"passwd"`   // 密码
	Code     string `toml:"code"`     // 命令
}
