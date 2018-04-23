using UnityEngine;

public class PersistBetweenScenes : MonoBehaviour
{
    private void Awake()
    {
        DontDestroyOnLoad(gameObject);
    }
}