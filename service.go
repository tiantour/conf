package conf

// Service service
type Service struct {
	Data service `toml:"service"`
}

// NewService new service
func NewService() *Service {
	return &conf.Service
}

type service struct {
	Account  string `toml:"account"`  // 帐号
	Article  string `toml:"article"`  // 文章
	Box      string `toml:"box"`      // 盒子
	Changer  string `toml:"changer"`  // 兑换
	Control  string `toml:"control"`  // 控制
	Discount string `toml:"discount"` // 折扣
	Favorite string `toml:"favorite"` // 收藏
	Goods    string `toml:"goods"`    // 商品
	Group    string `toml:"group"`    // 出团
	Lucky    string `toml:"lucky"`    // 抽奖
	Message  string `toml:"message"`  // 消息
	Monitor  string `toml:"monitor"`  // 监控
	Order    string `toml:"order"`    // 订单
	Paper    string `toml:"paper"`    // 问卷
	Product  string `toml:"product"`  // 产品
	Stock    string `toml:"stock"`    // 库存
	Ticket   string `toml:"ticket"`   // 票务
	Trade    string `toml:"trade"`    // 交易
	Vip      string `toml:"vip"`      // 会员
	Voyager  string `toml:"voyager"`  // 讲解
}
