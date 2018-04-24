using Logic;
using Model;
using UnityEngine;

namespace Utils {
    public class PlayerHud : MonoBehaviour {

        public GameBrain brain;
        private PlayerData me;

        private int verticalOffset;
        
        void Start() {
            verticalOffset = (int)(Screen.height / 15f);
            brain = FindObjectOfType<GameBrain>();
        }
        
        void Update() {
            if (brain != null) {
                if (me == null) {
                    // GetComponent is a costly function... this is terrible programming
                    var obj = brain.players.Find(p => p.GetComponent<PlayerData>().name == State.myName);
                    if (obj != null) me = obj.GetComponent<PlayerData>();
                } else {
                    var playerPos = me.transform.position;
                    var newScreenPos = Camera.main.WorldToScreenPoint(playerPos);
                    newScreenPos.y += verticalOffset;
//
                    newScreenPos.x -= Screen.width / 2f;
                    newScreenPos.y -= Screen.height / 2f;
                    Debug.Log("Arrow from world:" + playerPos + " to scren:" + newScreenPos);
                    transform.localPosition = newScreenPos;
//                    transform.localPosition = me.transform.position + new Vector3(0, 1, 0);
                }
            }
        }
    }
}