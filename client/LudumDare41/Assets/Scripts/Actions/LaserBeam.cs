using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LaserBeam : IActionScript
{
    public float time = 1.5f;
	public Transform parent;
	
	private GameObject _localLaserBeam;
    // Use this for initialization
    void Start ()
    {
	    _localLaserBeam = Instantiate(GetComponent<LaserBeamHolder>().LaserBeam, parent.position, parent.rotation);
	    soundPlayer.playSound(SoundsManager.SFX.TankFiring);
    }
	
	// Update is called once per frame
	void Update () {
        
        time -= Time.deltaTime;
		
//		transform.eulerAngles = transform.eulerAngles + new Vector3(0, rotation * Time.deltaTime, 0);
		if (time <= 0)
		{
			Destroy(_localLaserBeam);
			Destroy(this);
		}

	}
}
