using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class LobbyMember : MonoBehaviour {

	[HideInInspector] public Text name;
	// Use this for initialization
	void Start () {
		name = GetComponent<Text>();
	}
}
