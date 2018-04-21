using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;
using Utils;

public class SettingsBehaviour : MonoBehaviour {

	public InputField hostNameInput;

	void Start() {
		hostNameInput.text = State.host.Replace("http://", "").Replace("https://", "");
	}
	
	public void SaveHostName() {
		State.host = "http://" + hostNameInput.text;
		State.socketHost = "ws://" + hostNameInput.text;
		Back();
	}

	public void Back() {
		SceneManager.LoadScene("MainMenu");
	}
}
