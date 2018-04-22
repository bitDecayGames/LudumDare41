using System;
using System.Collections;
using System.Collections.Generic;
using Network.Messages;
using UnityEngine;
using Utils;

namespace Network {
    public class UpdateStream : MonoBehaviour {

        private static WebSocket webSocket;
        private static bool started = false;
        private List<IUpdateStreamSubscriber> subscribers = new List<IUpdateStreamSubscriber>();
        
        /// <summary>
        /// Requires you to wait for a bit before you can actually send messages
        /// </summary>
        public void StartListening(Action onSuccess) {
            if (!started) {
                StartCoroutine(startWebsocket(onSuccess));
            } else onSuccess.Invoke();
        }

        public void StopListening() {
            if (started) {
                started = false;
                webSocket.Close();
            }
        }

        public void Subscribe(IUpdateStreamSubscriber subscriber) {
            subscribers.Add(subscriber);
        }

        public void CancelSubscription(IUpdateStreamSubscriber subscriber) {
            if (subscribers.Contains(subscriber)) subscribers.Remove(subscriber);
        }

        public void Send(string msg) {
            if (started) webSocket.SendString(msg);
            else Debug.LogError("Failed to send message because Websocket was still initializing");
        }

        private IEnumerator startWebsocket(Action onSuccess) {
            Debug.Log("Attempt to connect to websocket: " + State.socketHost + "/api/v1/pubsub");
            var ws = new WebSocket(new Uri(State.socketHost + "/api/v1/pubsub"));
            yield return StartCoroutine(ws.Connect());
            started = true;
            webSocket = ws;
            onSuccess.Invoke();
            Debug.Log("Websocket now listening");
            while (started) {
                string msg = webSocket.RecvString();
                if (msg != null) {
                    var json = JsonUtility.FromJson<GenericUpdateStreamMessage>(msg);
                    subscribers.ForEach(s => s.receiveUpdateStreamMessage(json.messageType, msg));
                }

                if (webSocket.error != null) {
                    Debug.LogError("WebsocketError: " + webSocket.error);
                    break;
                }

                yield return 0;
            }

            StopListening();
        }
    }

    public interface IUpdateStreamSubscriber {
        void receiveUpdateStreamMessage(string messageType, string message);
    }
}