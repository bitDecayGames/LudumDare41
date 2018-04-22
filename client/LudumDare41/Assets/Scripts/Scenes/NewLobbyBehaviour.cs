using System.Collections;
using System.Collections.Generic;
using Network;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class NewLobbyBehaviour : MonoBehaviour {

	public InputField name;

	public void RequestNewLobby() {
		Debug.Log("Request new lobby");
		StartCoroutine(WebApi.RequestNewLobby((code) => {
			StartCoroutine(WebApi.JoinLobby(name.text, (lobby) => {
				SceneManager.LoadScene("Lobby");
			}, (err, status) => {
				Debug.LogError("Failed to join the newly created lobby(" + status + "): " + err);
			}));
		}, (err, status) => {
			Debug.LogError("Failed to create a new lobby(" + status + "): " + err);
		}));
	}
}
