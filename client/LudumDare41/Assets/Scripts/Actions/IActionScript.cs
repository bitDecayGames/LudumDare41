using Model;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public abstract class IActionScript : MonoBehaviour
{
    public Action<ActionData> onComplete;
    public ActionData actionData;

    void OnDestroy()
    {
        if(onComplete != null)
        {
            onComplete.Invoke(actionData);

        }
    }

}
