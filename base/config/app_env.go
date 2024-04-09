package config

type appEnvConfig struct {
	Env string `yaml:"app-env"`
}

var AppEnv = &appEnvConfig{
	Env: "DEV",
}
