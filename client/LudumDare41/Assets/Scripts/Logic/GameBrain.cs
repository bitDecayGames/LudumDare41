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
        public static int NUM_OF_CARDS_TO_SUBMIT = 3;

        public GameObject TilePrefab;
        public GameObject PlayerPrefab;
        public Hud HudPrefab;
        public EventSystem HudEventSystemPrefab;
        public List<TileMaterial> tileMaterials;
        public SoundsManager SoundPlayerPrefab;

        private Camera camera;
        private Hud hud;
        private SoundsManager SoundPlayer;

        private List<GameObject> tiles = new List<GameObject>();
        private List<GameObject> players = new List<GameObject>();
        private List<Step> stepSequence;
        bool actionInProgress;

        private ProcessedTurn currentTurn;
        private Action<List<Card>> userCardsSubmittion;
        private int actionsThisStep;
        private int stepsIndex;

        public bool isStepsComplete {
            get { return (stepsIndex >= currentTurn.diff.steps.Count); }
        }

        public bool isActionsComplete {
            get {
                return currentTurn != null &&
                       currentTurn.diff != null &&
                       currentTurn.diff.steps != null && 
                       stepsIndex < currentTurn.diff.steps.Count &&
                       currentTurn.diff.steps[stepsIndex].actions != null && 
                       actionsThisStep >= currentTurn.diff.steps[stepsIndex].actions.Count;
            }
        }


        private
            void Start() {
            camera = Camera.main;
            hud = Instantiate(HudPrefab);
            Instantiate(HudEventSystemPrefab);
            SoundPlayer = Instantiate(SoundPlayerPrefab);

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
//            if (Input.GetKeyDown(KeyCode.Space) ||
//                Input.GetKeyDown(KeyCode.KeypadEnter) ||
//                Input.GetKeyDown(KeyCode.Return) ||
//                Input.GetKeyDown(KeyCode.I)) {
//                ApplyTurn(TurnDebugger.GenerateTurn(), (s) => { s.ForEach(c => Debug.Log("C:" + c.id)); });
//            }

            if (isActionsComplete) stepCompleted();
        }

        /// <summary>
        /// Called by Network code to process each turn as it comes in from the server
        /// </summary>
        /// <param name="turn">the current processed turn</param>
        public void ApplyTurn(ProcessedTurn turn, Action<List<Card>> onSelected) {
            // based on turn start board, create the tile layout
            currentTurn = turn;
            userCardsSubmittion = onSelected;

            Debug.Log("start of apply turn");
            DestroyTiles();
            DestroyPlayers();
            GenerateTiles(turn.start.gameBoard.tiles);
            GeneratePlayers(turn.start.players);
            SetupCamera(turn.start.gameBoard.width);

            if (turn.diff != null && turn.diff.steps != null && turn.diff.steps.Count > 0) {
                stepsIndex = -1;
                stepCompleted();
            }
            else {
                applyEndofTurn();
            }
        }

        private void stepCompleted() {
            stepsIndex++;
            actionsThisStep = 0;
            if (!isStepsComplete) {
                var step = currentTurn.diff.steps[stepsIndex];

                foreach (ActionData action in step.actions) {
                    foreach (GameObject player in players) {
                        if (action.playerId == player.GetComponent<PlayerData>().name) {
                            IActionScript iAction = null;
                            switch (action.actionType) {
                                case "moveNorthAction":
                                    Move moveNComp = player.AddComponent<Move>();
                                    moveNComp.direction = new Vector3(0, 0, 1);
                                    iAction = moveNComp;
                                    break;
                                case "moveSouthAction":
                                    Move moveSComp = player.AddComponent<Move>();
                                    moveSComp.direction = new Vector3(0, 0, -1);
                                    iAction = moveSComp;
                                    break;
                                case "moveEastAction":
                                    Move moveEComp = player.AddComponent<Move>();
                                    moveEComp.direction = new Vector3(1, 0, 0);
                                    iAction = moveEComp;
                                    break;
                                case "moveWestAction":
                                    Move moveWComp = player.AddComponent<Move>();
                                    moveWComp.direction = new Vector3(-1, 0, 0);
                                    iAction = moveWComp;
                                    break;
                                case "rotateClockwiseAction":
                                    Rotate rotateCWComp = player.AddComponent<Rotate>();
                                    rotateCWComp.degrees = 90f;
                                    iAction = rotateCWComp;
                                    break;
                                case "rotateCounterClockwiseAction":
                                    Rotate rotateCCWComp = player.AddComponent<Rotate>();
                                    rotateCCWComp.degrees = -90f;
                                    iAction = rotateCCWComp;
                                    break;
                                case "rotate180Action":
                                    Rotate rotate180Comp = player.AddComponent<Rotate>();
                                    rotate180Comp.degrees = 180f;
                                    rotate180Comp.time = 2.5f;
                                    iAction = rotate180Comp;
                                    break;
                                case "shootMainGunAction":
                                    //player.AddComponent<Move>().direction = new Vector3(1, 0, 0);
                                    break;
                                case "deathAction":
                                    //player.AddComponent<Move>().direction = new Vector3(1, 0, 0);
                                    break;
                                case "spawnAction":
                                    Spawn spawnComp = player.AddComponent<Spawn>();
                                    iAction = spawnComp;
                                    break;
                            }

                            if (iAction != null) {
                                iAction.actionData = action;
                                iAction.onComplete = actionCompleted;
                            }
                        }
                    }
                }
            }
        }


        private void actionCompleted(ActionData data) {
            actionsThisStep++;
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

        private void applyEndofTurn() {
            DestroyTiles();
            DestroyPlayers();
            GenerateTiles(currentTurn.end.gameBoard.tiles);
            GeneratePlayers(currentTurn.end.players);
            var myPlayer = currentTurn.end.players.Find(p => p.name == State.myName);
            //if (myPlayer == null) myPlayer = turn.end.players[0]; // DEBUGGING ONLY
            if (myPlayer != null) {
                Debug.Log("Player: " + JsonUtility.ToJson(myPlayer, true));
                hud.ShowHand(myPlayer.hand, NUM_OF_CARDS_TO_SUBMIT, (selected) => {
                    userCardsSubmittion(selected);
                    hud.LowerCards();
                });
            }
            else {
                Debug.LogError("Failed to find my next hand of cards");
            }

            currentTurn = null;
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
                pos.y = 1;
                obj.transform.localPosition = pos;
                var pData = obj.GetComponent<PlayerData>();
                pData.name = p.name;
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