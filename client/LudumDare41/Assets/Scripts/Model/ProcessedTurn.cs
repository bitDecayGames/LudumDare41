using System.Collections.Generic;
using Model.Action.Abstract;

namespace Model {
    [System.Serializable]
    public class ProcessedTurn {
        public int tick;
        public GameState start;
        public GameState end;
        public List<Step> steps;
        public List<Card> inputs;
    }
}