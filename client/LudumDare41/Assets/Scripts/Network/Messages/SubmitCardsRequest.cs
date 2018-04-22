using System.Collections.Generic;
using Model;
using UnityEngine;

namespace Network.Messages {
    [System.Serializable]
    public class SubmitCardsRequest {
        public List<int> cardIds;

        public SubmitCardsRequest() {
            
        }

        public SubmitCardsRequest(List<Card> cards) {
            cards.ForEach(c => Debug.Log("Card: " + c.id));
            cardIds = cards.ConvertAll(c => c.id);
        }
    }
}