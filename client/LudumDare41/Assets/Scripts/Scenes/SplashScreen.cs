using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class SplashScreen : MonoBehaviour
{
    private const string FadeInCompletedStateName = "Fade-end";

    private Animator animator;

    private void Start() {
        animator = GetComponent<Animator>();
    }

    private void Update()
    {
        var animatorState = animator.GetCurrentAnimatorStateInfo(0);
        if (animatorState.IsName(FadeInCompletedStateName))
        {
            SceneManager.LoadScene("MainMenu");
        }
    }
}