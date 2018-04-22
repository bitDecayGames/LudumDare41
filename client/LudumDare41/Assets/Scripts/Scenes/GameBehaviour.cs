using Logic;
using Network;
using Network.Messages;
using UnityEngine;
using UnityEngine.SceneManagement;
using Utils;

namespace Scenes {
    public class GameBehaviour : MonoBehaviour, IUpdateStreamSubscriber {
        public GameBrain brain;

        void Start() {
            var updater = GetComponent<UpdateStream>();
            updater.Subscribe(this);
            updater.StartListening(() => { });
        }

        public void GetNewProcessedTurn() {
            StartCoroutine(WebApi.RefreshGameState((pt) => brain.ApplyTurn(pt, (selectedCards) => {
                StartCoroutine(WebApi.SubmitCardChoices(selectedCards, () => { }, (err, status) => Debug.LogError("Failed to submit card selections(" + status + "): " + err)));
            }), (err, status) => {
                Debug.LogError("Failed to apply turn(" + status + "): " + err);
            }));
        }

        public void receiveUpdateStreamMessage(string messageType, string message) {
            Debug.Log("Received update stream message:" + message);
            var json = JsonUtility.FromJson<GenericUpdateStreamMessage>(message);
            if (messageType == "connectionStarted") {
                Debug.Log("Saving connection id: " + json.id);
                State.connectionId = json.id;
                // get processed turn after connection id is reestablished
                StartCoroutine(WebApi.BroadcastConnectionId(GetNewProcessedTurn, (err, status) => {
                    Debug.LogError("Failed to broadcast connection id(" + status + "): " + err);
                }));
            } else if (messageType == "gameUpdate") {
                Debug.Log("Game is starting");
                State.currentTick = json.tick;
                GetNewProcessedTurn();
            }
        }
    }
}