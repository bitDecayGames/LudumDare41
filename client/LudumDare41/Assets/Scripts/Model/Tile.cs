using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Model;

namespace Model {
    [System.Serializable]
    public class Tile {
        public static string[] TILE_TYPES = new string[] {
            "empty",
            "wall"
        };
        
        public int id;
        public string tileType;
        public Vector pos;
    }
}