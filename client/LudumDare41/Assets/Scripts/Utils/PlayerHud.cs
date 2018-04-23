using Logic;
using UnityEngine;

namespace Utils {
    public class PlayerHud : MonoBehaviour {

        public GameBrain brain;
        private PlayerData me;

        private int verticalOffset = 30;
        
        void Start() {
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
                    transform.position = newScreenPos;
                }
            }
        }
    }
}