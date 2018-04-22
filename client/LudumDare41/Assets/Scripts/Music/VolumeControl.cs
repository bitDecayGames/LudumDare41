﻿using UnityEngine;

public class VolumeControl : MonoBehaviour
{
	private AudioSource _myAudioSource;

	private void Start()
	{
		GetComponent<AudioSource>().volume = Settings.MusicVolume;
	}
}