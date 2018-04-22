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
        public int x;
        public int y;
        public int z;
        Material DirtTile = Resources.Load("Materials/DirtTile", typeof(Material)) as Material;

        Material GrassTile = Resources.Load("Materials/GrassTile", typeof(Material)) as Material;

        Material FireTile = Resources.Load("Materials/FireTile", typeof(Material)) as Material;

        Material MountainTile = Resources.Load("Materials/MountainTile", typeof(Material)) as Material;

        Material WaterTile = Resources.Load("Materials/WaterTile", typeof(Material)) as Material;
        Material Coin = Resources.Load("Materials/Coin", typeof(Material)) as Material;

        GameObject tile = Resources.Load("Prefabs/Tile") as GameObject;


        public GameObject init()
        {

            switch (tileType)
            {
                case "Fire":                    
                    tile.GetComponentInChildren<MeshRenderer>().material = FireTile;
                    break;
                case "Water":
                    tile.GetComponentInChildren<MeshRenderer>().material = WaterTile;
                    break;
                case "Dirt":
                    tile.GetComponentInChildren<MeshRenderer>().material = DirtTile;
                    break;
                case "Grass":
                    tile.GetComponentInChildren<MeshRenderer>().material = GrassTile;
                    break;
                case "Mountain":
                    tile.GetComponentInChildren<MeshRenderer>().material = MountainTile;
                    break;
            }

            
            return GameObject.Instantiate(tile, new Vector3(x, y, z), Quaternion.identity);           
            
        }

        public Material getMaterial()
        {

            switch (tileType)
            {
                case "Fire":
                    return FireTile; 
                case "Water":
                    return WaterTile;
                case "Dirt":
                    return DirtTile;
                case "Grass":
                    return GrassTile;
                case "Mountain":
                    return MountainTile;
                default:
                    return Coin;
            }
        }

    }
}