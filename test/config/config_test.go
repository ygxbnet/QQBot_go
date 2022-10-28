package config

import (
	"QQBot_go/internal/config"
	"testing"
)

func TestParse(t *testing.T) {
	t.Log("Account.AdminID", config.Parse().Account.AdminID)
	t.Log("Account.BotID", config.Parse().Account.BotID)

	t.Log("Group.MainID", config.Parse().Group.MainID)
	t.Log("Group.InfoID", config.Parse().Group.InfoID)

	t.Log("Server.Websocket.URL", config.Parse().Server.Websocket.URL)
	t.Log("Server.HTTPAPI.URL", config.Parse().Server.HTTPAPI.URL)

	for i := 0; i <= 10000; i++ {
		config.Parse()
	}
}
