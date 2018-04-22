using System.Collections;
using System.Collections.Generic;
using Network;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class JoinLobbyBehaviour : MonoBehaviour {

	public InputField name;
	public InputField code;

	public void JoinLobby() {
		StartCoroutine(WebApi.JoinLobby(name.text, (lobby) => {
			SceneManager.LoadScene("Lobby");
		}, (err, status) => {
			Debug.LogError("Failed to join the newly created lobby(" + status + "): " + err);
		}));
	}
}
