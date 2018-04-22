using System.Collections.Generic;

namespace Network.Messages {
    [System.Serializable]
    public class LobbyPlayersResponse {
        public List<string> players;
    }
}