using System.Collections.Generic;
using Model;

namespace Network.Messages {
    [System.Serializable]
    public class SubmitCardsRequest {
        public List<int> cardIds;

        public SubmitCardsRequest() {
            
        }

        public SubmitCardsRequest(List<Card> cards) {
            cardIds = cards.ConvertAll(c => c.id);
        }
    }
}