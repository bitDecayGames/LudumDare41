using System;

namespace Model {
    [Serializable]
    public class ProcessedTurn {
        public int tick;
        public GameState start;
        public GameState end;
        public StepSequence diff;
    }
}