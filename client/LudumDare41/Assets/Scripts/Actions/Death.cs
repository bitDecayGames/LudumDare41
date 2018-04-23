using UnityEngine;

public class Death : IActionScript
{
	public float duration = 1.5f;
	private float speed;

	void Start () {
		speed = .33f / duration;
        soundPlayer.playSound(SoundsManager.SFX.TankDeath);
    }
	
	void Update () {
		duration -= Time.deltaTime;
		var scaleDelta = Time.deltaTime * speed;
		transform.localScale = transform.localScale - new Vector3(scaleDelta, scaleDelta, scaleDelta);
		if (duration <= 0) {
			GetComponentInChildren<SkinnedMeshRenderer>().enabled = false;
			Destroy(this);
		}
	}
}
