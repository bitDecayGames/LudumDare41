using System.Collections.Generic;

namespace Model {
    [System.Serializable]
    public class GameState {
        public int tick;
        public GameBoard gameBoard;
        public List<Player> players;
    }
}