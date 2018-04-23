using UnityEngine;

public class CanvasFadeIn : MonoBehaviour
{
    private CanvasGroup _canvasGroup;
    
    private void Start()
    {
        _canvasGroup = GetComponent<CanvasGroup>();
        if (_canvasGroup == null)
        {
            _canvasGroup = gameObject.AddComponent<CanvasGroup>();
        }
        _canvasGroup.alpha = 0f;
    }

    private void FixedUpdate()
    {
        if (_canvasGroup.alpha < 1f)
        {
            _canvasGroup.alpha += .01f;
        }
        else
        {
            Destroy(this);
        }
    }
}