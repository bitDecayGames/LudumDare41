using System;
using System.Collections.Generic;
using System.Linq;
using Model;
using UnityEngine;
using UnityEngine.UI;

namespace Prefabs {
    public class Hud : MonoBehaviour {
        public List<ButtonCard> Cards = new List<ButtonCard>();
        public List<CardImage> CardImages;
        private List<ButtonCard> selected = new List<ButtonCard>();
        private int numberToSelect = -1;
        private Action<List<Card>> onSelectionComplete;

        // HACK: this is probably bad...
        private static float DEFAULT = -20f;
        private static float LOWERED = -200f;
        private static float RAISED = 10f;
        private static float MAX_DIFF = 50f;

        private static float ANIM_SPEED = 0.1f;
        private float bottomOfScreenY = 0;
        
        void Start() {
            bottomOfScreenY = -Screen.height / 2f;
            
            var i = 0;
            Cards.ForEach(c => {
                var index = i;
                c.button.onClick.AddListener(() => {
                    SelectCard(index);
                });
                c.desiredHeight = LOWERED;
                var pos = c.transform.localPosition;
                pos.y = bottomOfScreenY + c.desiredHeight;
                c.transform.localPosition = pos;
                i++;
            });
            LowerCards();
        }

        void Update() {
            Cards.ForEach(c => {
                var pos = c.transform.localPosition;
                pos.y = pos.y + (bottomOfScreenY + c.desiredHeight - pos.y) * ANIM_SPEED;
                c.transform.localPosition = pos;
            });
        }
        
        public void ShowHand(List<Card> inputs, int numberToSelect, Action<List<Card>> onSelectionComplete) {
            if (numberToSelect <= 0) throw new Exception("Number to select cannot be less or equal to zero");
            if (numberToSelect > inputs.Count) throw new Exception("Number to select cannot be more than the number of cards");
            if (Cards.Count != inputs.Count) throw new Exception("Inputs count(" + inputs.Count + ") was different than Card count(" + Cards.Count + ")");
            
            selected.Clear();
            this.numberToSelect = numberToSelect;
            for (int i = 0; i < inputs.Count; i++) {
                var card = Cards[i];
                var data = inputs[i];
                var cardImage = GetCardImageForCardType(data.cardType);
                if (cardImage != null) card.button.image.sprite = cardImage.sprite;
                card.card = data;
                card.priorityText.text = "" + data.priority;
                card.debugText.text = data.cardType;
                card.selected = false;
                card.desiredHeight = DEFAULT;
                card.button.enabled = true;
            }
            this.onSelectionComplete = onSelectionComplete;
        }

        public void SelectCard(int index) {
            var card = Cards[index];
            card.selected = !card.selected;
            card.desiredHeight = DEFAULT;
            if (card.selected) selected.Add(card);
            else selected.Remove(card);
            for (int i = 0; i < selected.Count; i++) {
                var sel = selected[i];
                sel.desiredHeight = RAISED + (1f - ((float)i / selected.Count)) * MAX_DIFF;
            }

            if (selected != null && onSelectionComplete != null && selected.Count >= numberToSelect) {
                onSelectionComplete.Invoke(selected.ConvertAll(c => c.card));
            }
        }

        public void LowerCards() {
            selected.Clear();
            Cards.ForEach(c => {
                c.desiredHeight = LOWERED;
                c.selected = false;
                c.card = null;
                c.priorityText.text = "";
                c.debugText.text = "";
                c.button.enabled = false;
            });
        }

        private CardImage GetCardImageForCardType(string cardType) {
            return CardImages.Find(c => c.cardType == cardType);
        }
    }

    [System.Serializable]
    public class CardImage {
        public string cardType;
        public Sprite sprite;
    }
}