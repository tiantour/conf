package iot

// Scene scene
type Scene struct {
	Name        string   // 名称
	Phone       []string // 电话
	PUE         string   // pue
	CPU         string   // 芯片
	System      string   // 系统
	Temperature string   // 温度
	Humidity    string   // 湿度
	Noise       string   // 噪音
	Import      []string // 进口
	Export      []string // 出口
	Pressure    []string // 压力
	Difference  []string // 差值
}
