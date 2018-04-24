using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Spawn : IActionScript
{	
	public float duration = 2.5f;
	private bool inited = false;
	private float curY = 0;

	private static float initialAltitude = 1;
	private static float fallingAltitude = 10;
	private static float speed = 0.95f;
	
	void Update() {
		Vector3 pos = transform.localPosition;
		if (!inited && actionData != null && actionData.position != null) {
			if (pos != null) {
				inited = true;
				pos.x = actionData.position.x;
				pos.z = actionData.position.y;
				pos.y = fallingAltitude;
				curY = fallingAltitude;
				transform.localPosition = pos;
				GetComponentInChildren<SkinnedMeshRenderer>().enabled = true;
				transform.localScale = new Vector3(0.33f, 0.33f, 0.33f);
			}
		}

		if (inited) {
			duration -= Time.deltaTime;
			curY *= speed;
			pos.y = curY + initialAltitude;
			
			transform.localPosition = pos;
			if (duration <= 0) {
				Destroy(this);
			}
		}
	}
}
