package conf

// Gotour gotour
type Gotour struct {
	Data gotour `toml:"gotour"`
}

// NewGotour new gotour
func NewGotour() *Gotour {
	return &conf.Gotour
}

// gotour gotour
type gotour struct {
	Account string `toml:"account"` // 账户
	Article string `toml:"article"` // 文章
	Box     string `toml:"box"`     // 盒子
	Cattle  int32  `toml:"cattle"`  // 黄牛
	Control string `toml:"control"` // 监控
	Lucky   string `toml:"lucky"`   // 抽奖
	Paper   string `toml:"paper"`   // 问卷
	Trade   string `toml:"trade"`   // 交易
}
