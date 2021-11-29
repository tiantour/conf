package conf

// Scene scene
type Scene struct {
	Data map[string]scene `toml:"scene"`
}

// NewScene new scene
func NewScene() *Scene {
	return &conf.Scene
}

type scene struct {
	Name        string   `json:"name"`        // 名称
	Phone       []string `toml:"phone"`       // 电话
	PUE         string   `toml:"pue"`         // pue
	CPU         string   `toml:"cpu"`         // 芯片
	System      string   `toml:"system"`      // 系统
	Temperature string   `toml:"temperature"` // 温度
	Humidity    string   `toml:"humidity"`    // 湿度
	Noise       string   `toml:"noise"`       // 噪音
	Import      []string `toml:"import"`      // 进口
	Export      []string `toml:"export"`      // 出口
	Pressure    []string `toml:"pressure"`    // 压力
	Difference  []string `toml:"difference"`  // 差值
}
