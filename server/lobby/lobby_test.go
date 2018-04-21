package lobby

import "testing"

func TestLobbyPlayerAdd(t *testing.T) {
	lobbyService := NewLobbyService()

	lobby, err := lobbyService.NewLobby()
	if err != nil {
		t.Fatal(err)
	}

	if len(lobbyService.GetLobbies()) != 1 {
		t.Error("Lobby not added to the service")
	}

	if len(lobbyService.GetLobbies()[0].Name) < 6 {
		t.Error("lobby name is too short")
	}

	err = lobby.AddPlayer("Jacque")
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
