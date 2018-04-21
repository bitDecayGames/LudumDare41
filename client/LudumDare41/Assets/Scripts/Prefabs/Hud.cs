using System;
using System.Collections.Generic;
using System.Linq;
using Model;
using UnityEngine;
using UnityEngine.UI;

namespace Prefabs {
    public class Hud : MonoBehaviour {
        public List<Button> ButtonObjects;

        private List<ButtonCard> cards = new List<ButtonCard>();
        private List<ButtonCard> selected = new List<ButtonCard>();
        private int numberToSelect = -1;
        private Action<List<Card>> onSelectionComplete;

        // HACK: this is probably bad...
        private static float DEFAULT = -230f;
        private static float LOWERED = -450;
        private static float RAISED = -210;
        private static float MAX_DIFF = 40f;

        private static float ANIM_SPEED = 0.1f;
        
        void Start() {
            var i = 0;
            ButtonObjects.ForEach(b => {
                var c = new ButtonCard(b, null);
                var index = i;
                b.onClick.AddListener(() => {
                    SelectCard(index);
                });
                cards.Add(c);
                i++;
            });
            LowerCards();
        }

        void Update() {
            cards.ForEach(c => {
                var pos = c.button.transform.localPosition;
                pos.y = pos.y + (c.desiredHeight - pos.y) * ANIM_SPEED;
                c.button.transform.localPosition = pos;
            });
        }
        
        public void ShowHand(List<Card> inputs, int numberToSelect, Action<List<Card>> onSelectionComplete) {
            if (numberToSelect <= 0) throw new Exception("Number to select cannot be less or equal to zero");
            if (numberToSelect > inputs.Count) throw new Exception("Number to select cannot be more than the number of cards");
            if (cards.Count != inputs.Count) throw new Exception("Inputs count(" + inputs.Count + ") was different than Card count(" + cards.Count + ")");
            
            selected.Clear();
            this.numberToSelect = numberToSelect;
            for (int i = 0; i < inputs.Count; i++) {
                var card = cards[i];
                var data = inputs[i];
                card.card = data;
                card.text.text = data.id + ": " + data.cardType + " " + data.priority;
                card.selected = false;
                card.desiredHeight = DEFAULT;
                card.button.enabled = true;
            }
            this.onSelectionComplete = onSelectionComplete;
        }

        public void SelectCard(int index) {
            var card = cards[index];
            card.selected = !card.selected;
            card.desiredHeight = DEFAULT;
            if (card.selected) selected.Add(card);
            else selected.Remove(card);
            for (int i = 0; i < selected.Count; i++) {
                var sel = selected[i];
                sel.desiredHeight = RAISED + (1f - (float)(i + 1 / selected.Count)) * MAX_DIFF;
            }

            if (selected != null && onSelectionComplete != null && selected.Count >= numberToSelect) {
                onSelectionComplete.Invoke(selected.ConvertAll(c => c.card));
            }
        }

        public void LowerCards() {
            selected.Clear();
            cards.ForEach(c => {
                c.desiredHeight = LOWERED;
                c.selected = false;
                c.card = null;
                c.text.text = "";
                c.button.enabled = false;
            });
        }
    }

    public class ButtonCard {
        public Button button;
        public Text text;
        public Card card;
        public bool selected = false;
        public float desiredHeight = 0;

        public ButtonCard(Button button, Card card) {
            this.button = button;
            this.card = card;
            text = button.GetComponentInChildren<Text>();
        }
    }
}