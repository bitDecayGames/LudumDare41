﻿using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class MainMenuBehaviour : MonoBehaviour {

	public void NewLobby() {
		SceneManager.LoadScene("NewLobby");
	}

	public void JoinLobby() {
		SceneManager.LoadScene("JoinLobby");
	}

	public void Settings() {
		SceneManager.LoadScene("Settings");
	}
}
