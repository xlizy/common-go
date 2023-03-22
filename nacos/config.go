package nacos

var BaseWebConfigVal *BaseWebConfig

type WebServerConfig struct {
	Port         int    `yaml:"port"`
	CookieDomain string `yaml:"cookie-domain"`
}
type BaseWebConfig struct {
	ServerConfig WebServerConfig   `yaml:"server"`
	AppEnv       string            `yaml:"app-env"`  //发布环境
	AppSign      map[string]string `yaml:"app-sign"` //应用间秘钥
	AppName      string            //应用名
}
