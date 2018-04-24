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
        public GameObject PlayerHudPrefab;
        public Hud HudPrefab;
        public EventSystem HudEventSystemPrefab;
        public List<TileMaterial> tileMaterials;
        public SoundsManager SoundPlayerPrefab;
        public MapFactory MapFactoryPrefab;

        private Camera camera;
        private Hud hud;
        private SoundsManager SoundPlayer;
        private MapFactory map;

        private List<GameObject> tiles = new List<GameObject>();
        public List<GameObject> players = new List<GameObject>();
        private List<Step> stepSequence;
        bool actionInProgress;

        private ProcessedTurn currentTurn;
        private Action<List<Card>> userCardsSubmittion;
        private int actionsThisStep;
        private int stepsIndex;
        private bool turnOffTiles = true;
        private GameObject playerHud;

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


        private void Start() {
            camera = Camera.main;
            hud = Instantiate(HudPrefab);
            Instantiate(HudEventSystemPrefab);
            map = Instantiate(MapFactoryPrefab, transform);
            map.transform.localPosition = new Vector3(0, 0, 0);
            SoundPlayer = Instantiate(SoundPlayerPrefab);
            var canvas = hud.GetComponent<Canvas>();
            canvas.renderMode = RenderMode.ScreenSpaceCamera;
            canvas.worldCamera = camera;
            
//            SoundPlayer.playSoundLoop(SoundsManager.SFX.EngineIdleLoop);


            var mapSkin = map.BuildMap("1stMap");
            if (mapSkin != null) mapSkin.transform.localPosition = new Vector3(0, 0.99f, 0);
        }


        private void SetupCamera(int boardWidth) {
            var pos = camera.transform.position;
            pos.x = transform.position.x + boardWidth / 2f - 0.5f;
            pos.y = transform.position.y + 8;
            pos.z = transform.position.z - 4.5f;
            camera.transform.position = pos;
            var canvas = hud.GetComponent<Canvas>();
            canvas.worldCamera = camera;
            canvas.planeDistance = 3f;

            camera.transform.eulerAngles = new Vector3(55, 0, 0);
        }

        void Update() {
            // DEBUG
//            if (Input.GetKeyDown(KeyCode.Space) ||
//                Input.GetKeyDown(KeyCode.KeypadEnter) ||
//                Input.GetKeyDown(KeyCode.Return) ||
//                Input.GetKeyDown(KeyCode.I))
//            {
//                ApplyTurn(TurnDebugger.GenerateTurn(), (s) => { s.ForEach(c => Debug.Log("C:" + c.id)); });
//            }
//
//            if (Input.GetKeyDown(KeyCode.T)) {
//                turnOffTiles = !turnOffTiles;
//                if (turnOffTiles) DestroyTiles();
//                else if (currentTurn != null) GenerateTiles(currentTurn.start.gameBoard.tiles);
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

            if (!turnOffTiles) {
                DestroyTiles();
                GenerateTiles(turn.start.gameBoard.tiles);
            }
            DestroyPlayers();
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
                Debug.Log("Actions for this step: " + JsonUtility.ToJson(step, true));
                foreach (ActionData action in step.actions) {
                    if(action.playerId == "gameBoard")
                    {
                        IActionScript iAction = null;
                        switch (action.actionType)
                        {
                            case "setNextCrateAction":
                            default:
                                Debug.LogError("Failed to handle action: " + action.actionType);
                                actionCompleted(action);
                                break;
                        }
                        if (iAction != null)
                        {
                            iAction.actionData = action;
                            iAction.onComplete = actionCompleted;
                        }

                    }
                    else foreach (GameObject player in players) {
                        if (action.playerId == player.GetComponent<PlayerData>().name) {
                            IActionScript iAction = null;
                            switch (action.actionType.ToUpper()) {
                                case "MOVENORTHACTION":
                                    Move moveNComp = player.AddComponent<Move>();
                                    moveNComp.direction = new Vector3(0, 0, 1);
                                    moveNComp.soundPlayer = SoundPlayer;
                                    iAction = moveNComp;
                                    break;
                                case "MOVESOUTHACTION":
                                    Move moveSComp = player.AddComponent<Move>();
                                    moveSComp.direction = new Vector3(0, 0, -1);
                                    moveSComp.soundPlayer = SoundPlayer;
                                    iAction = moveSComp;
                                    break;
                                case "MOVEEASTACTION":
                                    Move moveEComp = player.AddComponent<Move>();
                                    moveEComp.direction = new Vector3(1, 0, 0);
                                    moveEComp.soundPlayer = SoundPlayer;
                                    iAction = moveEComp;
                                    break;
                                case "MOVEWESTACTION":
                                    Move moveWComp = player.AddComponent<Move>();
                                    moveWComp.direction = new Vector3(-1, 0, 0);
                                    moveWComp.soundPlayer = SoundPlayer;
                                    iAction = moveWComp;
                                    break;
                                case "ROTATECLOCKWISEACTION":
                                    Rotate rotateCWComp = player.AddComponent<Rotate>();
                                    rotateCWComp.degrees = 90f;
                                    rotateCWComp.soundPlayer = SoundPlayer;
                                    iAction = rotateCWComp;
                                    break;
                                case "ROTATECOUNTERCLOCKWISEACTION":
                                    Rotate rotateCCWComp = player.AddComponent<Rotate>();
                                    rotateCCWComp.degrees = -90f;
                                    rotateCCWComp.soundPlayer = SoundPlayer;
                                    iAction = rotateCCWComp;
                                    break;
                                case "ROTATE180ACTION":
                                    Rotate rotate180Comp = player.AddComponent<Rotate>();
                                    rotate180Comp.degrees = 180f;
                                    rotate180Comp.time = 2.5f;
                                    rotate180Comp.soundPlayer = SoundPlayer;
                                    iAction = rotate180Comp;
                                    break;
                                case "SPAWNACTION":
                                    Spawn spawnComp = player.AddComponent<Spawn>();
                                    spawnComp.soundPlayer = SoundPlayer;
                                    iAction = spawnComp;
                                    break;
                                case "DEATHACTION":
                                    Death deathComp = player.AddComponent<Death>();
                                    deathComp.soundPlayer = SoundPlayer;
                                    iAction = deathComp;
                                    break;
                                case "SHOOTMAINGUNACTION":
                                    Debug.Log("laserbeam");
                                    LaserBeam laserComp = player.AddComponent<LaserBeam>();
                                    laserComp.soundPlayer = SoundPlayer;
                                    laserComp.parent = player.transform;
                                    iAction = laserComp;
                                    break;
                                default:
                                    Debug.LogError("Failed to handle action: " + action.actionType);
                                    actionCompleted(action);
                                break;
                            }

                            if (iAction != null) {
                                iAction.actionData = action;
                                iAction.onComplete = actionCompleted;
                                iAction.tankAnimation = player.GetComponentInChildren<AnimateTank>();
                            }
                        }
                    }
                }
            }
            else
            {
                applyEndofTurn();
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
                pos.y = 1;
                obj.transform.localPosition = pos;
                var mesh = obj.GetComponentInChildren<MeshRenderer>();
                mesh.material = tileMaterials.Find(m => m.name == t.tileType).material;
            });
        }

        private void applyEndofTurn() {
            if (!turnOffTiles) {
                DestroyTiles();
                GenerateTiles(currentTurn.end.gameBoard.tiles);
            }

            DestroyPlayers();
            GeneratePlayers(currentTurn.end.players);
            var myPlayer = currentTurn.end.players.Find(p => p.name == State.myName);
            //if (myPlayer == null) myPlayer = turn.end.players[0]; // DEBUGGING ONLY
            if (myPlayer != null) {
                Debug.Log("End of turn Player state: " + JsonUtility.ToJson(myPlayer, true));
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
                var yrot = 0f;
                if (p.facing.x > 0) yrot = 180;
                else if (p.facing.x < 0) yrot = 0;
                else if (p.facing.y > 0) yrot = 90;
                else if (p.facing.y < 0) yrot = -90;
                obj.transform.eulerAngles = new Vector3(0,yrot,0);

                if (p.name == State.myName) {
                    playerHud = Instantiate(PlayerHudPrefab, obj.transform);
                    playerHud.transform.localPosition = new Vector3(0, 2, 0);
                }
                
            });
        }

        private void DestroyPlayers() {
            if (playerHud != null) {
                Destroy(playerHud);
                playerHud = null;
            }
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