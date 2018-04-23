using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Rotate : MonoBehaviour {
    public string rotateType;
    private float endDegrees;
    private float startingDegrees;
	// Use this for initialization
	void Start () {
        startingDegrees = transform.rotation.y;
        endDegrees = startingDegrees - 45;
	}
	
	// Update is called once per frame
	void Update () {

        float rotation = Time.deltaTime;
        float translation = Time.deltaTime * 10;
        transform.Translate(0, 0, translation);
        Debug.Log("MOVING");
        if (duration <= 0)
            Destroy(this);
    }
}
