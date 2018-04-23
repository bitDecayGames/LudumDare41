using UnityEngine;

public class LoopAfter : MonoBehaviour
{
	public AudioSource FirstAudioSource;
	private AudioSource _myAudioSource;
	private float _startTime;
	
	private void Start()
	{
		_myAudioSource = GetComponent<AudioSource>();
		_startTime = Time.time;
	}

	void Update() {
		Debug.Log("Current time: " + Time.time);
		if (Time.time - _startTime > 3f && FirstAudioSource != null && _myAudioSource != null)
		{
			if (!FirstAudioSource.isPlaying && !_myAudioSource.isPlaying)
			{
				_myAudioSource.Play();
				_myAudioSource.loop = true;
			}		
		}
	}
}