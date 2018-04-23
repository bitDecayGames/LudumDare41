using Model;
using Network;
using UnityEngine;
using Utils;

namespace Scenes {
	public class MwDebugBehaviour : MonoBehaviour , IUpdateStreamSubscriber {

		void Start () {
			var a = TurnDebugger.GenerateTurn();
			var str = JsonUtility.ToJson(a);
			var b = JsonUtility.FromJson<ProcessedTurn>(str);
			
			Debug.Log(str);
			Debug.Log(JsonUtility.ToJson(b));
		}

		public void receiveUpdateStreamMessage(string messageType, string message) {
			
		}
	}
}
