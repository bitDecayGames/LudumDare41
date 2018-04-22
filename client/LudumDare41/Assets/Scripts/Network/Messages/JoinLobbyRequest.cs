namespace Network.Messages {
    [System.Serializable]
    public class JoinLobbyRequest {
        public string playerName;

        public JoinLobbyRequest() {
            
        }

        public JoinLobbyRequest(string playerName) {
            this.playerName = playerName;
        }
    }
}