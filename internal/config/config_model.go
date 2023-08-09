package config

type Config struct {
	PrependMessage string
	Account        Account
	Group          Group
	Server         Server
	OpenAI         OpenAI
}

type Account struct {
	BotID   string
	AdminID string
}

type Group struct {
	MainID string
	InfoID string
}

type Server struct {
	HttpAPI   HttpAPI
	Websocket Websocket
}
type Websocket struct {
	URL string
}
type HttpAPI struct {
	URL string
}

type OpenAI struct {
	BaseURL string
	APIKey  []string
}
