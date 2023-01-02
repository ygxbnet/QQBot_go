package config

type Config struct {
	PrependMessage string  `yaml:"prepend-message"`
	Account        Account `yaml:"account"`
	Group          Group   `yaml:"group"`
	Server         Server  `yaml:"server"`
}
type Account struct {
	BotID   string `yaml:"bot-id"`
	AdminID string `yaml:"admin-id"`
}
type Group struct {
	MainID string `yaml:"main-id"`
	InfoID string `yaml:"info-id"`
}
type Websocket struct {
	URL string `yaml:"url"`
}
type HTTPAPI struct {
	URL string `yaml:"url"`
}
type Server struct {
	HTTPAPI   HTTPAPI   `yaml:"http-api"`
	Websocket Websocket `yaml:"websocket"`
}
