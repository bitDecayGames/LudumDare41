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
        public float raisedHeight {
            get { return originalHeight + 10; }
        }
        [HideInInspector]
        public float loweredHeight {
            get { return originalHeight - 600; }
        }
        [HideInInspector]
        public float originalHeight = 0;


        void Awake() {
            button = GetComponent<Button>();
        }

        public void Init(Card card) {
            this.card = card;
        }

        public void SetHeightToRaised() {
            var pos = button.transform.localPosition;
            pos.y = raisedHeight;
            button.transform.localPosition = pos;
        }

        public void SetHeightToLowered() {
            var pos = button.transform.localPosition;
            pos.y = loweredHeight;
            button.transform.localPosition = pos;
        }

        public void SetHeightToOriginal() {
            Debug.Log("Set height to original");
            var pos = button.transform.localPosition;
            pos.y = originalHeight;
            button.transform.localPosition = pos;
        }
    }
}