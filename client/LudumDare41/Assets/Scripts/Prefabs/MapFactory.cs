using System.Collections.Generic;
using UnityEngine;

namespace Prefabs {
    public class MapFactory : MonoBehaviour {
        public List<Map> Maps;
    }

    [System.Serializable]
    public class Map {
        public string name;
        public GameObject map;
    }
}