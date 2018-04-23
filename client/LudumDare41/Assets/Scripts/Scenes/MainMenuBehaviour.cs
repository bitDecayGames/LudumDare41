using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class MainMenuBehaviour : MonoBehaviour
{
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
