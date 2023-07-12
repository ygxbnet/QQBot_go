package config

import (
	"QQBot_go/internal/config"
	"testing"
)

func TestParse(t *testing.T) {
	t.Log("Account.AdminID", config.Get().Account.AdminID)
	t.Log("Account.BotID", config.Get().Account.BotID)

	t.Log("Group.MainID", config.Get().Group.MainID)
	t.Log("Group.InfoID", config.Get().Group.InfoID)

	t.Log("Server.Websocket.URL", config.Get().Server.Websocket.URL)
	t.Log("Server.HTTPAPI.URL", config.Get().Server.HTTPAPI.URL)

	for i := 0; i <= 10000; i++ {
		config.Get()
	}
}
