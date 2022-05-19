package config

var C_Log map[string]string = make(map[string]string)

func init() {
	RegisterParser("log.ymal", &C_Log)
}
