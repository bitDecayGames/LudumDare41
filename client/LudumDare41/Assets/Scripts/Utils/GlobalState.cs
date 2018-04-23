using UnityEngine;

public class GlobalState : MonoBehaviour
{
    private static GlobalState instance;

    public bool HasFadedIn;
    
    public static GlobalState Instance
    {
        get
        {
            if (instance == null)
            {
                GameObject gameObject = new GameObject();
                instance = gameObject.AddComponent<GlobalState>();
                gameObject.name = "GlobalState";
            }

            return instance;
        }
    }

    void Awake()
    {
        if (instance == null)
        {
            instance = this;
        }
        
        DontDestroyOnLoad(gameObject);
    }
}