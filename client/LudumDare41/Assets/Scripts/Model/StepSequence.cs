using System;
using System.Collections.Generic;

namespace Model {
    [Serializable]
    public class StepSequence {
        public List<Card> cards;
        public List<Step> steps;
    }
}