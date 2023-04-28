package config

type appSignConfig struct {
	Sign map[string]string `yaml:"sign"`
}

var AppSign = &appSignConfig{
	Sign: make(map[string]string),
}
