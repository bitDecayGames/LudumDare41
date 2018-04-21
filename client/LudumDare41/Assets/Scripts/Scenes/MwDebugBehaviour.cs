using Network;
using UnityEngine;

namespace Scenes {
	public class MwDebugBehaviour : MonoBehaviour {

		// Use this for initialization
		void Start () {
			
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
		}
	}
}
