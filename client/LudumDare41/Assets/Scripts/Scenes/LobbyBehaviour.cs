using Network;
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
		isFirstPlayer = State.lobby.players.FindIndex(p => p.name == State.myName) == 0;
		
		startBtn.gameObject.SetActive(isFirstPlayer);
		waitingText.gameObject.SetActive(!isFirstPlayer);
		
		RefreshLobbyMembers();

		if (isFirstPlayer) startBtn.enabled = false;
		GetComponent<UpdateStream>().StartListening(() => { startBtn.enabled = true; });
	}

	public void RefreshLobbyMembers() {
		foreach (Transform child in lobbyMembers) {
			GameObject.Destroy(child.gameObject); // remove all lobby member prefabs
		}
		State.lobby.players.ForEach(p => {
			var lobbyMember = Instantiate(lobbyMemberPrefab);
			lobbyMember.transform.SetParent(lobbyMembers);
			lobbyMember.name.text = p.name;
		});
	}

	public void StartGame() {
		// TODO: send a start game request
	}

	public void receiveUpdateStreamMessage(string message) {
		// TODO: if message is something like: "RefreshLobbyMembers" then refresh
		// TODO: if message is something like: "RequestTick0" then move to game board?
	}
}
