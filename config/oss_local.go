package config

type Local struct {
	Url         string `mapstructure:"url" json:"url" yaml:"url"`                            // 本地文件访问路径
	Path        string `mapstructure:"path" json:"path" yaml:"path"`                         // 本地文件访问路径
	StorePath   string `mapstructure:"store-path" json:"store-path" yaml:"store-path"`       // 本地文件存储路径
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 上传文件预命名缓存时间（秒）
}
