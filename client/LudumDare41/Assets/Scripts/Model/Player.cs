using System.Collections.Generic;

namespace Model {
    [System.Serializable]
    public class Player {
        public int id;
        public string name;
        public string color;
        public int x;
        public int y;
        public int facingX;
        public int facingY;
        public List<Card> hand;
    }
}