package conf

// Cache cache
type Cache struct {
	Data cache `toml:"cache"`
}

// NewCache new cache
func NewCache() *Cache {
	return &conf.Cache
}

type cache struct {
	IP   string `toml:"ip"`   // 地址
	Port string `toml:"port"` // 端口
}
