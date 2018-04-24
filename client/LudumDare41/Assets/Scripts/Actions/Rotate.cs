using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Rotate : IActionScript
{
    public float degrees;
    public float time = 1.5f;
    private float rotation;
    // Use this for initialization
    void Start () {
        rotation = degrees / time;
        soundPlayer.playSound(SoundsManager.SFX.EngineRev); GetComponentInChildren<AnimateTank>().enabled = true;


    }
	
	// Update is called once per frame
	void Update () {
        
        time -= Time.deltaTime;
        transform.eulerAngles = transform.eulerAngles + new Vector3(0, rotation * Time.deltaTime, 0);
        if (time <= 0)
            Destroy(this);
    }
}
