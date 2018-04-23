using System.Collections.Generic;

namespace Model {
    [System.Serializable]
    public class Player {
        public string name;
        public string color;
        public int team;
        public Vector pos;
        public Vector facing;
        public List<Card> hand;
    }
}