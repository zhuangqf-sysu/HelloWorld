package config

// config 信息
type Config struct {
	path string
	//配置信息
	// db、redis、es、自定义配置项
}

func NewConfig(path string) (*Config, error) {
	return &Config{path: path}, nil
}

// todo 从yaml文件加载
// todo 动态加载：定时拉取
