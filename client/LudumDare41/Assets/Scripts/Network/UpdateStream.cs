using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Network {
    public class UpdateStream : MonoBehaviour {
        public static string host = "ws://localhost:8080/api/v1/pubsub";
//        public static string host = "ws://echo.websocket.org"; // for debugging websockets

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
            var ws = new WebSocket(new Uri(host));
            yield return StartCoroutine(ws.Connect());
            started = true;
            webSocket = ws;
            onSuccess.Invoke();
            Debug.Log("Websocket now listening");
            while (started) {
                string msg = webSocket.RecvString();
                if (msg != null) {
                    subscribers.ForEach(s => s.receiveUpdateStreamMessage(msg));
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
        void receiveUpdateStreamMessage(string message);
    }
}