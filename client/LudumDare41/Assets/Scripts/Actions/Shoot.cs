using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Shoot : IActionScript {
    //just an animation


	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
        // Play the shoot animation, k?
        Destroy(this);
	}
}
