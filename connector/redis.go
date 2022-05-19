package connector

type RedisConf struct {
	Addr           string `yaml:"Addr"`
	MaxActive      int    `yaml:"MaxActive"`
	MaxIdle        int    `yaml:"MaxIdle"`
	IdleTimeout    int    `yaml:"IdleTimeout"`
	ConnectTimeout int    `yaml:"ConnectTimeout"`
	ReadTimeout    int    `yaml:"ReadTimeout"`
	WriteTimeout   int    `yaml:"WriteTimeout"`
	TestInterval   int    `yaml:"TestInterval"`
}
