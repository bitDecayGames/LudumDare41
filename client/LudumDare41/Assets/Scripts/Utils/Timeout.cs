using UnityEngine;
using UnityEngine.Events;

namespace Utils {
    public class Timeout : MonoBehaviour {
        public float delay = 0; // defaults to no delay
        public float cycle = 10; // default to 10 seconds
        public UnityEvent onTimeout;

        private int count = 0;
        private float currentTime = 0;

        void Update() {
            currentTime += Time.deltaTime;
            if (count == 0 && currentTime >= delay) TriggerTimeout();
            else if (currentTime >= cycle) TriggerTimeout();
        }

        void TriggerTimeout() {
            count++;
            currentTime = 0;
            onTimeout.Invoke();
        }
    }
}