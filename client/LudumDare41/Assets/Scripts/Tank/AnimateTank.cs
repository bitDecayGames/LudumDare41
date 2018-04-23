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
            if (m.name == "Material_005")
            {
                m.mainTextureOffset = new Vector2(OffsetX, OffsetY);
            }
        }
    }
}