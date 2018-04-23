namespace Model {
    [System.Serializable]
    public class Card {
        public static string[] CARD_TYPES = new string[] {
            "moveForward1Card",
            "moveForward2Card",
            "moveForward3Card",
            "moveBackwardCard",
            "rotateClockwiseCard",
            "rotateCounterclockwiseCard",
            "rotate180Card",
            "shootMainTurretCard"
        };
            
            
        public int id;
        public string cardType;
        public int priority;
        public string owner;
    }
}