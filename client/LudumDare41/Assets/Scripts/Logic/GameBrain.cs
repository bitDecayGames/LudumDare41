using System;
using System.Collections.Generic;
using Model;
using Network;
using Prefabs;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;
using Utils;

namespace Logic {
    public class GameBrain : MonoBehaviour {

        public GameObject TilePrefab;
        public GameObject PlayerPrefab;
        public Hud HudPrefab;
        public EventSystem HudEventSystemPrefab;
        public List<TileMaterial> tileMaterials;

        private Camera camera;
        private Hud hud;

        private List<GameObject> tiles = new List<GameObject>();
        private List<GameObject> players = new List<GameObject>();

        void Start() {
            Debug.Log("Start game brain");
            camera = Camera.main;
            hud = Instantiate(HudPrefab);
            Instantiate(HudEventSystemPrefab);
        }

        private void SetupCamera(int boardWidth) {
            var pos = camera.transform.position;
            pos.x = transform.position.x + boardWidth / 2f - 0.5f;
            pos.y = transform.position.y + 7;
            pos.z = transform.position.z - 4;
            camera.transform.position = pos;
            var canvas = hud.GetComponent<Canvas>();
            canvas.worldCamera = camera;
            canvas.planeDistance = 3f;
                
            
            camera.transform.eulerAngles = new Vector3(55, 0, 0);
        }

        void Update() {
            if (Input.GetKeyDown(KeyCode.Space) ||
                Input.GetKeyDown(KeyCode.KeypadEnter) ||
                Input.GetKeyDown(KeyCode.Return) || 
                Input.GetKeyDown(KeyCode.I)) {
                ApplyTurn(TurnDebugger.GenerateTurn());
            }
        }

        /// <summary>
        /// Called by Network code to process each turn as it comes in from the server
        /// </summary>
        /// <param name="turn">the current processed turn</param>
        public void ApplyTurn(ProcessedTurn turn) {
            // based on turn start board, create the tile layout
            DestroyTiles();
            DestroyPlayers();
            GenerateTiles(turn.start.board.tiles);
            GeneratePlayers(turn.start.players);
            SetupCamera(turn.start.board.width);

            // TODO: based on the turn steps, create sequences of actions
            // TODO: all of these methods will eventually need to become asynchronous to handle the animation delays
            
            // TODO: based on turn end board, recreate the tile layout
            // DestroyTiles();
            // DestroyPlayers();
            //GenerateTiles(turn.end.board.tiles); // TODO: uncomment this
            //GeneratePlayers(turn.end.players); // TODO: uncomment this
            //var nextHand = turn.end.players.Find(p => p.name == State.myName).hand; // TODO: uncomment this
            var nextHand = turn.end.players[0].hand;
            hud.ShowHand(nextHand, 3, (selected) => {
                // TODO: uncomment this
//                WebApi.SubmitCardChoices(selected, () => { }, (err, status) => {
//                    Debug.LogError("Failed to send card choices(" + status + "): " + err);
//                });
                hud.LowerCards();
            });
        }

        private void GenerateTiles(List<Tile> tileData) {
            tileData.ForEach(t => {
                var obj = Instantiate(TilePrefab, transform);
                tiles.Add(obj);
                var pos = obj.transform.localPosition;
                pos.x = t.x;
                pos.z = t.y;
                pos.y = 0;
                obj.transform.localPosition = pos;
                var mesh = obj.GetComponentInChildren<MeshRenderer>();
                mesh.material = tileMaterials.Find(m => m.name == t.tileType).material;
            });
        }

        private void DestroyTiles() {
            tiles.ForEach(Destroy);
            tiles.Clear();
        }

        private void GeneratePlayers(List<Player> playerData) {
            playerData.ForEach(p => {
                var obj = Instantiate(PlayerPrefab, transform);
                players.Add(obj);
                var pos = obj.transform.localPosition;
                pos.x = p.x;
                pos.z = p.y;
                pos.y = 0;
                obj.transform.localPosition = pos;
            });
        }

        private void DestroyPlayers() {
            players.ForEach(Destroy);
            players.Clear();
        }
    }

    [Serializable]
    public class TileMaterial {
        public string name;
        public Material material;
    }
}
