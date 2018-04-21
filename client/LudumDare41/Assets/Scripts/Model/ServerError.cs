namespace Model {
    [System.Serializable]
    public class ServerError {
        #region Required

        public string message;

        #endregion

        #region Optional

        public int tick = -1;

        #endregion
    }
}