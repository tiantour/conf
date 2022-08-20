package iot

// Scene scene
type Scene struct {
	Name        string   `mapstructure:"name"`        // 名称
	Phone       []string `mapstructure:"phone"`       // 电话
	PUE         string   `mapstructure:"pue"`         // pue
	CPU         string   `mapstructure:"cpu"`         // 芯片
	System      string   `mapstructure:"system"`      // 系统
	Temperature string   `mapstructure:"temperature"` // 温度
	Humidity    string   `mapstructure:"humidity"`    // 湿度
	Noise       string   `mapstructure:"noise"`       // 噪音
	Import      []string `mapstructure:"import"`      // 进口
	Export      []string `mapstructure:"export"`      // 出口
	Pressure    []string `mapstructure:"pressure"`    // 压力
	Difference  []string `mapstructure:"difference"`  // 差值
}
