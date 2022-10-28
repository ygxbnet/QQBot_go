package config

type Config struct {
	Account Account `yaml:"account"`
	Group   Group   `yaml:"group"`
	Server  Server  `yaml:"server"`
}
type Account struct {
	BotID   int64 `yaml:"bot-id"`
	AdminID int64 `yaml:"admin-id"`
}
type Group struct {
	MainID int `yaml:"main-id"`
	InfoID int `yaml:"info-id"`
}
type Websocket struct {
	URL string `yaml:"url"`
}
type HTTPAPI struct {
	URL string `yaml:"url"`
}
type Server struct {
	Websocket Websocket `yaml:"websocket"`
	HTTPAPI   HTTPAPI   `yaml:"http-api"`
}
