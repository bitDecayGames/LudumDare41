﻿using System;
using Network;
using Network.Messages;
using UnityEngine;
using UnityEngine.UI;
using Utils;

public class LobbyBehaviour : MonoBehaviour, IUpdateStreamSubscriber {

	public Text title;
	public Text waitingText;
	public Transform lobbyMembers;
	public LobbyMember lobbyMemberPrefab;
	public Button startBtn;

	private bool isFirstPlayer = false;
	
	void Start () {
		if (State.lobby == null) throw new Exception("Lobby was null in the Lobby Scene... whoops...");

		title.text = "Lobby: " + State.lobby.code;
		
		isFirstPlayer = State.lobby.players.FindIndex(p => p.name == State.myName) == 0;
		
		startBtn.gameObject.SetActive(isFirstPlayer);
		waitingText.gameObject.SetActive(!isFirstPlayer);
		
		RefreshLobbyMembers();

		if (isFirstPlayer) startBtn.enabled = false;
		var updater = GetComponent<UpdateStream>();
		updater.Subscribe(this);
		updater.StartListening(() => { startBtn.enabled = true; });
	}

	public void RefreshLobbyMembers() {
		foreach (Transform child in lobbyMembers) {
			Destroy(child.gameObject); // remove all lobby member prefabs
		}
		State.lobby.players.ForEach(p => {
			var lobbyMember = Instantiate(lobbyMemberPrefab);
			lobbyMember.transform.SetParent(lobbyMembers);
			lobbyMember.name.text = p.name;
		});
	}

	public void StartGame() {
		Debug.Log("Trying to start the game");
		// TODO: send a start game request
	}

	public void receiveUpdateStreamMessage(string messageType, string message) {
		Debug.Log("Received update stream message:" + message);
		var json = JsonUtility.FromJson<GenericUpdateStreamMessage>(message);
		if (messageType == "connectionStarted") {
			Debug.Log("Saving connection id: " + json.id);
			State.connectionId = json.id;
			StartCoroutine(WebApi.BroadcastConnectionId(() => { }, (err, status) => {
				Debug.LogError("Failed to broadcast connection id(" + status + "): " + err);
			}));
		}
		if (messageType == "playerJoin") {
			Debug.Log("Player joined: " + json.id);
			StartCoroutine(WebApi.RefreshCurrentLobby((l) => {
				RefreshLobbyMembers();
			}, (err, status) => {
				Debug.LogError("Failed to refresh lobby id(" + status + "): " + err);
			}));
		}
		// TODO: if message is something like: "RefreshLobbyMembers" then refresh
		// TODO: if message is something like: "RequestTick0" then move to game board?
	}
}