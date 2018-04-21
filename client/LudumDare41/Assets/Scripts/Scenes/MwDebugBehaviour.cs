using Network;
using UnityEngine;
using Utils;

namespace Scenes {
	public class MwDebugBehaviour : MonoBehaviour , IUpdateStreamSubscriber {

		// Use this for initialization
		void Start () {
			var updater = GetComponent<UpdateStream>();
			updater.StartListening(() => {});
			updater.Subscribe(this);
		}
	
		// Update is called once per frame
		void Update () {
		
		}

		public void CheckServer() {
			Debug.Log("Attempting Ping...");
			StartCoroutine(WebApi.Ping(() => {
				Debug.Log("Server is up and running");
			}, (msg, status) => {
				Debug.Log("Error response(" + status + "): " + msg);
			}));
			
			GetComponent<UpdateStream>().Send("Hello ping");
		}

		public void receiveUpdateStreamMessage(string messageType, string message) {
			Debug.Log("Successfully subscribed and got message: " + message);
		}
	}
}
