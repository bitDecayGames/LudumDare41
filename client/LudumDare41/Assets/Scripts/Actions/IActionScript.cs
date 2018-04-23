using Model;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public abstract class IActionScript : MonoBehaviour
{
    public Action<ActionData> onComplete;
    public ActionData actionData;
    public SoundsManager soundPlayer;

    void OnDestroy()
    {
        if(onComplete != null)
        {
            soundPlayer.playSoundLoop(SoundsManager.SFX.EngineIdleLoop);
            onComplete.Invoke(actionData);

        }
    }

}
