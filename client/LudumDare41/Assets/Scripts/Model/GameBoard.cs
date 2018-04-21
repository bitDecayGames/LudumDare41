using System.Collections.Generic;

namespace Model {
    [System.Serializable]
    public class GameBoard {
        public int width;
        public int height;
        public List<Tile> tiles;
    }
}