package lobby

import "testing"

func TestLobbyPlayerAdd(t *testing.T) {
	lobbyService := NewLobbyService()

	lobby := lobbyService.NewLobby()

	err := lobby.AddPlayer("Jacque")
	if err != nil {
		t.Error(err)
	}
	err = lobby.AddPlayer("Monday")
	if err != nil {
		t.Error(err)
	}

	err = lobby.AddPlayer("Monday")
	if err == nil {
		t.Error("Allowed duplicate player names")
	}
}
