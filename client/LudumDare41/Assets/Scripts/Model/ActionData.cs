using System;

namespace Model {
    [Serializable]
    public class ActionData {
        public static string[] ACTION_TYPES = new string[] {
            "moveNorthAction",
            "moveSouthAction",
            "moveEastAction",
            "moveWestAction",
            "rotateClockwiseAction",
            "rotateCounterClockwiseAction",
            "rotate180Action",
            "shootMainGunAction",
            "deathAction",
            "spawnAction"
        };
        public int id;
        public string playerId;
        public string actionType;
        public Vector position;
    }
}