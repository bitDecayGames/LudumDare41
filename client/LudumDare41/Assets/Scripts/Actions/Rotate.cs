using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Rotate : MonoBehaviour {
    public string rotateType;
    public float degrees;
	// Use this for initialization
	void Start () {
	}
	
	// Update is called once per frame
	void Update () {

        float rotation = Time.deltaTime;
        degrees -= Time.deltaTime;
        transform.Rotate(0, rotation, 0);
        Debug.Log("Rotating");
        if (degrees <= 0)
            Destroy(this);
    }
}
