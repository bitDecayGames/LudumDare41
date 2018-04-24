using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Move : IActionScript
{
    public float duration = 1.5f;
    public float distance = 1f;
    private float speed;
    public Vector3 direction;

    
    // Use this for initialization
    void Start () {
        speed = distance / duration;
        soundPlayer.playSound(SoundsManager.SFX.EngineRev);
        GetComponentInChildren<AnimateTank>().enabled = true;

    }
	
	// Update is called once per frame
	void Update () {

        duration -= Time.deltaTime;
        transform.position = transform.position + (direction * Time.deltaTime * speed);
        if (duration <= 0)
            Destroy(this);
    }
}
