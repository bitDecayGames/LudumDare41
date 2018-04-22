using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Model;

namespace Model {
    [System.Serializable]
    public class Tile
    {
        public int id;
        public string tileType;
        public Vector pos;
    }
}