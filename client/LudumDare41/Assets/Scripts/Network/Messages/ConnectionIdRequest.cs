namespace Network.Messages {
    [System.Serializable]
    public class ConnectionIdRequest {
        public string gameName;
        public string playerName;

        public ConnectionIdRequest() {
            
        }

        public ConnectionIdRequest(string gameName, string playerName) {
            this.gameName = gameName;
            this.playerName = playerName;
        }
    }
}