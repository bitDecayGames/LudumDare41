using System;
using UnityEngine;

public class SceneFadeIn : MonoBehaviour
{
    public Texture2D fadeOutTexture;

    private const int DirectionFadeIn = -1;
    private const int DirectionFadeOut = 1;
    private const float PreFadeIn = 1f;
    private const float PreFadeOut = 0f;

    private float fadeSpeed = 0.2f;
    private float maxFadeAmountPerFrame = .005f;
    private int drawDepth = -1000;

    private int _fadeDir;
    private float _alpha;


    private void Start()
    {
        FadeIn();
    }

    void OnGUI()
    {
        _alpha += _fadeDir * fadeSpeed * Time.deltaTime;
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
        _fadeDir = DirectionFadeIn;
        if (!GlobalState.Instance.HasFadedIn)
        {
            _alpha = PreFadeIn;
        }
        else
        {
            _alpha = PreFadeOut;
        }
    }

    public void FadeOut()
    {
        _fadeDir = DirectionFadeOut;
        _alpha = 0f;
    }
}