package lobby

import "testing"

func TestLobbyPlayerAdd(t *testing.T) {
	lobbyService := NewLobbyService()

	lobby := lobbyService.NewLobby()

	if len(lobbyService.GetLobbies()) != 1 {
		t.Error("Lobby not added to the service")
	}

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
