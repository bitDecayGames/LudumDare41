using UnityEngine;

public class PlayAfter : MonoBehaviour
{
	public AudioSource FirstAudioSource;
	private AudioSource _myAudioSource;

	private void Start()
	{
		_myAudioSource = GetComponent<AudioSource>();
	}

	void Update() {
		if (Time.time > 3f && FirstAudioSource != null && _myAudioSource != null)
		{
			if (!FirstAudioSource.isPlaying && !_myAudioSource.isPlaying)
			{
				_myAudioSource.Play();
			}		
		}
	}
}