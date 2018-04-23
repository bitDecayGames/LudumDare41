using UnityEngine;

public class AnimateTank : MonoBehaviour
{

    public float ScrollX = 0.25f;
    public float ScrollY = 0.25f;

    private void Update()
    {
        float OffsetX = Time.time * ScrollX;
        float OffsetY = Time.time * ScrollY;
        var materials = GetComponent<Renderer>().materials;
        foreach (Material m in materials)
        {
            Debug.Log("Material name: " + m.name);
            if (m.name == "TreadAnimation (Instance)")
            {
                Debug.Log("Found the right material!");
                m.mainTextureOffset = new Vector2(OffsetX, OffsetY);
            }
        }
    }
}