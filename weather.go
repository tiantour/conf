package conf

// Weather weather
type Weather struct {
	Data weather `toml:"weather"`
}

// NewWeather new weather
func NewWeather() *Weather {
	return &conf.Weather
}

type weather struct {
	ID  string `toml:"id"`  // 编号
	Key string `toml:"key"` // 键值
}
