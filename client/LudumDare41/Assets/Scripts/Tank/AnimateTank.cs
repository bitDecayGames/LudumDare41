using UnityEngine;

public class AnimateTank : MonoBehaviour {
    public float ScrollX = 0.25f;
    public float ScrollY = 0.25f;

    private Material correctMaterial;

    void Start() {
        var materials = GetComponent<Renderer>().materials;
        foreach (Material m in materials) {
            //Debug.Log("Material name: " + m.name);
            if (m.name == "TreadAnimation (Instance)") {
                //Debug.Log("Found the right material!");
                correctMaterial = m;
                break;
            }
        }
    }

    private void Update() {
        if (correctMaterial != null) {
            float OffsetX = Time.time * ScrollX;
            float OffsetY = Time.time * ScrollY;
            correctMaterial.mainTextureOffset = new Vector2(OffsetX, OffsetY);
        }
    }
}