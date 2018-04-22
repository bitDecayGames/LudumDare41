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
		Debug.Log("Attempt to join lobby(" + code.text + ") as " + name.text);
		StartCoroutine(WebApi.JoinLobby(name.text, code.text, (lobby) => {
			Debug.Log("Joined lobby: " + lobby);
			StartCoroutine(WebApi.RefreshCurrentLobby((newestLobby) => {
				SceneManager.LoadScene("Lobby");
			}, (err, status) => {
				Debug.LogError("Failed to get newest lobby(" + status + "): " + err);
			}));
		}, (err, status) => {
			Debug.LogError("Failed to join the newly created lobby(" + status + "): " + err);
		}));
	}
}
