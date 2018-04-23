using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Move : MonoBehaviour {
    public float duration;
	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {

        duration -= Time.deltaTime;
        float translation = Time.deltaTime * 10;
        transform.Translate(0, 0, translation);

        if (duration <= 0)
            Destroy(this);
    }
}
