﻿using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Death : IActionScript
{

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {
        GetComponentInChildren<SkinnedMeshRenderer>().enabled = false;
        Destroy(this);
	}
}
