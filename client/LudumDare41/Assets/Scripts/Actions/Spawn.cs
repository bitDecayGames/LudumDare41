using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Spawn : IActionScript
{

    
    //Todo spawning at end location

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
		if (actionData != null && actionData.position != null) {
			Vector3 pos = transform.position;
			if (pos != null) {
				pos.x = actionData.position.x;
				pos.y = actionData.position.y;
				transform.position = pos;
				GetComponentInChildren<SkinnedMeshRenderer>().enabled = true;
				Destroy(this);
			}
		}
	}
}
