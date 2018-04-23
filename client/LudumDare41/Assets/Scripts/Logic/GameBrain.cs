using System;
using System.Collections.Generic;
using Model;
using Model.Action.Abstract;
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
        public GameObject SoundPlayerObj;
        SoundsManager SoundPlayer;

        private Camera camera;
        private Hud hud;

        private List<GameObject> tiles = new List<GameObject>();
        private List<GameObject> players = new List<GameObject>();
        private List<Step> stepSequence;
        bool actionInProgress;
        public bool hasSteps
        {
            get { return (stepSequence.Count > 0); }
        }


        private 

        void Start() {
            camera = Camera.main;
            hud = Instantiate(HudPrefab);
            Instantiate(HudEventSystemPrefab);

            SoundPlayer = SoundPlayerObj.GetComponent<SoundsManager>();
            SoundPlayer.playSound(SoundsManager.SFX.TankFiring);
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
                ApplyTurn(TurnDebugger.GenerateTurn(), (s) => {
                    s.ForEach(c => Debug.Log("C:" + c.id));
                });
            }

            if (actionInProgress)
                return;

            
             
            
        }

        /// <summary>
        /// Called by Network code to process each turn as it comes in from the server
        /// </summary>
        /// <param name="turn">the current processed turn</param>
        public void ApplyTurn(ProcessedTurn turn, Action<List<Card>> onSelected) {
            // based on turn start board, create the tile layout
            DestroyTiles();
            DestroyPlayers();
            GenerateTiles(turn.start.gameBoard.tiles);
            GeneratePlayers(turn.start.players);
            SetupCamera(turn.start.gameBoard.width);
            players[0].AddComponent<Move>().duration = 5;
            stepSequence = turn.steps;
            // TODO: based on the turn steps, create sequences of actions
            // TODO: all of these methods will eventually need to become asynchronous to handle the animation delays

            // based on turn end board, recreate the tile layout
            DestroyTiles();
            DestroyPlayers();
            GenerateTiles(turn.end.gameBoard.tiles);
            GeneratePlayers(turn.end.players);
            var myPlayer = turn.end.players.Find(p => p.name == State.myName);
            //if (myPlayer == null) myPlayer = turn.end.players[0]; // DEBUGGING ONLY
            if (myPlayer != null) {
                Debug.Log("Player: " + JsonUtility.ToJson(myPlayer, true));
                hud.ShowHand(myPlayer.hand, 3, (selected) => {
                    onSelected(selected);
                    hud.LowerCards();
                });
            } else {
                Debug.LogError("Failed to find my next hand of cards");
            }
        }

        private void GenerateTiles(List<Tile> tileData) {
            tileData.ForEach(t => {
                var obj = Instantiate(TilePrefab, transform);
                tiles.Add(obj);
                var pos = obj.transform.localPosition;
                pos.x = t.pos.x;
                pos.z = t.pos.y;
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
                pos.x = p.pos.x;
                pos.z = p.pos.y;
                pos.y = 0;
                obj.transform.localPosition = pos;
                var pData = obj.GetComponent<PlayerData>();
                pData.id= p.id;

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
