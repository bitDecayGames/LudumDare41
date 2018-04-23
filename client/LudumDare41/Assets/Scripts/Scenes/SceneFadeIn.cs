using System;
using UnityEngine;

public class SceneFadeIn : MonoBehaviour
{
    public Texture2D fadeOutTexture;
    private const float PreFadeInAlpha = 1f;

    private float fadeSpeed = 0.2f;
    private float maxFadeAmountPerFrame = .005f;
    private int drawDepth = -1000;

    private int _fadeDir = -1;
    private float _alpha;


    private void Start()
    {
        FadeIn();
    }

    void OnGUI()
    {
        _alpha += _fadeDir * fadeSpeed * Time.deltaTime;
        Debug.Log("Current alpha: " + _alpha);
        _alpha = Mathf.Clamp(_alpha, _alpha - maxFadeAmountPerFrame, _alpha + maxFadeAmountPerFrame);
        _alpha = Mathf.Clamp01(_alpha);

        GUI.color = new Color(GUI.color.r, GUI.color.g, GUI.color.b, _alpha);
        GUI.depth = drawDepth;
        GUI.DrawTexture(new Rect(0, 0, Screen.width, Screen.height), fadeOutTexture);

        if (Math.Abs(_alpha - 1) < .05f)
        {
            GlobalState.Instance.HasFadedIn = true;
        }
    }

    public void FadeIn()
    {
        if (!GlobalState.Instance.HasFadedIn)
        {
            _alpha = PreFadeInAlpha;
        }
        else
        {
            _alpha = 0f;
        }
    }
}