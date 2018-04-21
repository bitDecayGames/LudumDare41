using System.Collections.Generic;

namespace Model {
    [System.Serializable]
    public class GameState {
        public int tick;
        public GameBoard board;
        public List<Player> players;
    }
}