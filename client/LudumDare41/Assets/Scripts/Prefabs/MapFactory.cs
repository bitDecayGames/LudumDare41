using System.Collections.Generic;
using UnityEngine;

namespace Prefabs {
    public class MapFactory : MonoBehaviour {
        public List<Map> Maps;

        public GameObject BuildMap(string mapName) {
            var mapPrefab = Maps.Find(m => m.name == mapName);
            if (mapPrefab != null) return Instantiate(mapPrefab.map, transform);
            return null;
        }
    }

    [System.Serializable]
    public class Map {
        public string name;
        public GameObject map;
    }
}