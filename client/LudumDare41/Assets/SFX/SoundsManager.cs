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
    void Start () {
        audioSource = GetComponent<AudioSource>();

    }
	
	// Update is called once per frame
	void Update () {
		
	}

    public void playSound(SFX sound){
        foreach(AudioClip clip in SoundsList)
        {
            if (clip.name == sound.ToString())
            {                
                audioSource.clip = clip;
                audioSource.Play();
            }
        }
        
    }

}
