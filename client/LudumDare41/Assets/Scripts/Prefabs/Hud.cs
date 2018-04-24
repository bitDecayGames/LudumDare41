using System;
using System.Collections.Generic;
using System.Linq;
using Model;
using UnityEngine;
using UnityEngine.UI;
using Utils;

namespace Prefabs {
    public class Hud : MonoBehaviour {
        public List<ButtonCard> Cards = new List<ButtonCard>();
        public List<CardImage> CardImages;
        private List<ButtonCard> selected = new List<ButtonCard>();
        private int numberToSelect = -1;
        private Action<List<Card>> onSelectionComplete;

        private static float ANIM_SPEED = 0.1f;

        void Start() {
            var i = 0;
            Cards.ForEach(c => {
                var index = i;
                c.button.onClick.AddListener(() => {
                    SelectCard(index);
                });
                c.originalHeight = c.button.transform.localPosition.y;
                i++;
            });
            LowerCards();
        }

        void Update() {
//            Cards.ForEach(c => {
//                var pos = c.transform.localPosition;
//                pos.y = pos.y + (bottomOfScreenY + c.desiredHeight - pos.y) * ANIM_SPEED;
//                c.transform.localPosition = pos;
//            });
            
            if (Input.GetKeyDown(KeyCode.A)) {
                var cards = new List<Card>();
                for (int i = 0; i < 5; i++) cards.Add(TurnDebugger.GenerateCard());
                ShowHand(cards, 3, (c)=> {
                });
            }
            else if (Input.GetKeyDown(KeyCode.S)) {
                LowerCards();
            }
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
                card.button.enabled = true;
                card.SetHeightToOriginal();
            }
            this.onSelectionComplete = onSelectionComplete;
        }

        public void SelectCard(int index) {
            var card = Cards[index];
            card.selected = !card.selected;
            if (card.selected) {
                selected.Add(card);
                card.SetHeightToRaised();
            }
            else {
                selected.Remove(card);
                card.SetHeightToOriginal();
            }
            for (int i = 0; i < selected.Count; i++) {
                var sel = selected[i];
            }

            if (selected != null && onSelectionComplete != null && selected.Count >= numberToSelect) {
                onSelectionComplete.Invoke(selected.ConvertAll(c => c.card));
            }
        }

        public void LowerCards() {
            selected.Clear();
            Cards.ForEach(c => {
                c.SetHeightToLowered();
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