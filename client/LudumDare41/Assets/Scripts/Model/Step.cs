using System.Collections.Generic;
using Model.Action.Abstract;

namespace Model {
    [System.Serializable]
    public class Step {
        public string id;
        public List<IAction> actions;
    }
}