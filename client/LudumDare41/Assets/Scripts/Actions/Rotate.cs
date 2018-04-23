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
	}
	
	// Update is called once per frame
	void Update () {
        
        time -= Time.deltaTime;
        transform.Rotate(0, rotation*Time.deltaTime, 0);
        Debug.Log("Rotating");
        if (time <= 0)
            Destroy(this);
    }
}
