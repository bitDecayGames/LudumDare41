using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SoundsManager : MonoBehaviour {


    public List<AudioClip> SoundsList;
    AudioSource audioSource;

    public enum SFX
    {
        TankFiring,
        TankEngineLoop,
        EngineIdleLoop,
        EngineRev,
        CrateOpen,
        TankDeath
    }
    // Use this for initialization
    void Awake () {
        audioSource = GetComponent<AudioSource>();
        Debug.Log("Audio source: " + audioSource);
    }
	
	// Update is called once per frame
	void Update () {
		
	}

    public void playSound(SFX sound){
        foreach(AudioClip clip in SoundsList)
        {
            if (clip.name == sound.ToString())
            {
                Debug.Log("clip name: " + clip.name);
                Debug.Log("sound enum: " + sound);
                audioSource.clip = clip;
                audioSource.Play();
            }
        }
        
    }

}
