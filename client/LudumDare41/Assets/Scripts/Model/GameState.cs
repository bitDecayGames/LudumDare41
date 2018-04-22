namespace Model {
    [System.Serializable]
    public class GameState {
        public int tick;
        public GameBoard board;

        public GameState()
        {
            board = new GameBoard();
        }
    }
}