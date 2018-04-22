using Model;

namespace Utils {
    public static class State {
        public static string host = "http://localhost:8080";
        public static string socketHost = "ws://localhost:8080";
        public static Player me = null;
        public static Lobby lobby = null;
        public static string myName = null;
        public static GameState state = null;
        public static ProcessedTurn processedTurn = null;
        public static int currentTick = 0;
        public static string connectionId = null;
    }
}