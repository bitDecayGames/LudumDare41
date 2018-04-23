using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Shoot : IActionScript {
    //just an animation
    private float duration = 1.0f;

	// Use this for initialization
	void Start ()
    {
        soundPlayer.playSound(SoundsManager.SFX.TankFiring);

    }
	
	// Update is called once per frame
	void Update () {
        // Play the shoot animation, k?   duration -= Time.deltaTime;
        duration -= Time.deltaTime;
        if (duration <= 0)
            Destroy(this);
	}
}
