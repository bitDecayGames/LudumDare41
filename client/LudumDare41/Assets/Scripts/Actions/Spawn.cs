using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Spawn : IActionScript
{	
	public float duration = 1.5f;
	private float speed;
	private bool inited = false;

	void Start() {
		speed = 1 / duration;
	}
	
	void Update() {
		if (!inited && actionData != null && actionData.position != null) {
			Vector3 pos = transform.position;
			if (pos != null) {
				inited = true;
				pos.x = actionData.position.x;
				pos.y = actionData.position.y;
				transform.position = pos;
				transform.localScale = new Vector3(0, 0, 0);
				GetComponentInChildren<SkinnedMeshRenderer>().enabled = true;
			}
		}

		duration -= Time.deltaTime;
		var scaleDelta = Time.deltaTime * speed;
		transform.localScale = transform.localScale + new Vector3(scaleDelta, scaleDelta, scaleDelta);
		if (duration <= 0) {
			Destroy(this);
		}
	}
}
