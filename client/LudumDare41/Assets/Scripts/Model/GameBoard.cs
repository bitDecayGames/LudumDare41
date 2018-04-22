using System;
using System.Collections.Generic;
using UnityEngine;

namespace Model {
    [System.Serializable]
    public class GameBoard : MonoBehaviour {
        public int width;
        public int height;
        public List<Tile> Tiles;
        public List<GameObject> Gametiles = new List<GameObject>();

        public GameBoard()
        {
            Tiles = new List<Tile>();
        }

        private void Start()
        { 
        }

        public void Update()
        {
        }

        public void init()
        {
            foreach (Tile tile in Tiles)
            {
                Gametiles.Add(tile.init());
            }
        }

        public void refresh()
        {
            foreach (GameObject gameTile in Gametiles)
            {
                foreach(Tile serverTile in Tiles)
                {
                    Vector3 pos = new Vector3(serverTile.x, serverTile.y, serverTile.z);
                    if (pos == gameTile.GetComponent<Transform>().transform.position)
                    {
                        gameTile.GetComponentInChildren<MeshRenderer>().material = serverTile.getMaterial();
                    }
                }
            }
        }

    }
}