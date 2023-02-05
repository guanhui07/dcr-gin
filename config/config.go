package config

// DataSourceConfig 数据库的结构体
type DataSourceConfig struct {
	DriverName string `mapstructure:"driverName"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Charset    string `mapstructure:"charset"`
	Loc        string `mapstructure:"loc"`
}

// ServerConfig 整个项目的配置
type ServerConfig struct {
	Port         int              `mapstructure:"port"`
	DataSource   DataSourceConfig `mapstructure:"datasource"`
	Local        Local            `mapstructure:"local" json:"local" yaml:"local"`
	Redis        Redis            `mapstructure:"redis" json:"redis" yaml:"redis"`
	Hashids      Hashids          `mapstructure:"hashids" json:"hashids" yaml:"hashids"`
	UploadTicket UploadTicket     `mapstructure:"uploadTicket"`
}
type Hashids struct {
	Salt      string `mapstructure:"salt"`
	MinLength int    `mapstructure:"minLength"`
}
type UploadTicket struct {
	Url string `mapstructure:"url"`
}
