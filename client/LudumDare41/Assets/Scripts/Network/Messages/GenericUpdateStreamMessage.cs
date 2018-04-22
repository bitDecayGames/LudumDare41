namespace Network.Messages {
    [System.Serializable]
    public class GenericUpdateStreamMessage {
        public string messageType;
        public string id;
        public int tick;
    }
}