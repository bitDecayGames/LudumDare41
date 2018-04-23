using Model;
using UnityEngine;
using UnityEngine.UI;

namespace Prefabs {
    public class ButtonCard : MonoBehaviour {
        public Text priorityText;
        public Text debugText;
        [HideInInspector] 
        public Button button;
        [HideInInspector]
        public Card card;
        [HideInInspector]
        public bool selected = false;
        [HideInInspector]
        public float desiredHeight = 0;

        void Awake() {
            button = GetComponent<Button>();
        }

        public void Init(Card card) {
            this.card = card;
        }
    }
}