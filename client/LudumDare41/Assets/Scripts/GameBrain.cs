using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Model;
using Model.Action.Abstract;
using System;

public class GameBrain : MonoBehaviour {

    public GameBoard Board;
    public Canvas Hud;
    //Player Player1;
    //Player Player2;
    //Player Player3;
    //Player Player4;
    //GameObject[] Cards;

    ProcessedTurn turn = new ProcessedTurn();

    // Use this for initialization
    void Start()
    {

        turn.start = new GameState();
        turn.start.tick = 1;
        turn.start.board.width = 3; 
        turn.start.board.height = 3;
        turn.start.board.Tiles = new List<Tile>();
        
        turn.end = new GameState();
        turn.end.board.width = 3;
        turn.end.board.height = 3;
        turn.end.board.Tiles = new List<Tile>();

        Tile tile1 = new Tile();
        Tile tile2 = new Tile();
        Tile tile3 = new Tile();
        Tile tile4 = new Tile();
        Tile tile5 = new Tile();
        Tile tile6 = new Tile();
        Tile tile7 = new Tile();
        Tile tile8 = new Tile();
        Tile tile9 = new Tile();


        tile1.id = 1;
        tile1.tileType = "Fire";
        tile1.x = 0;
        tile1.y = 0;
        tile1.z = 0;

        tile2.id = 2;
        tile2.tileType = "Water";
        tile2.x = 0;
        tile2.y = 0;
        tile2.z = 1;

        tile3.id = 3;
        tile3.tileType = "Grass";
        tile3.x = 0;
        tile3.y = 0;
        tile3.z = 2;

        tile4.id = 4;
        tile4.tileType = "Mountain";
        tile4.x = 1;
        tile4.y = 0;
        tile4.z = 0;

        tile5.id = 5;
        tile5.tileType = "Dirt";
        tile5.x = 1;
        tile5.y = 0;
        tile5.z = 1;

        tile6.id = 6;
        tile6.tileType = "Grass";
        tile6.x = 1;
        tile6.y = 0;
        tile6.z = 2;

        tile7.id = 7;
        tile7.tileType = "Dirt";
        tile7.x = 2;
        tile7.y = 0;
        tile7.z = 0;

        tile8.id = 8;
        tile8.tileType = "Mountain";
        tile8.x = 2;
        tile8.y = 0;
        tile8.z = 1;

        tile9.id = 9;
        tile9.tileType = "Water";
        tile9.x = 2;
        tile9.y = 0;
        tile9.z = 2;

        turn.start.board.Tiles.Add(tile1);
        turn.start.board.Tiles.Add(tile2);
        turn.start.board.Tiles.Add(tile3);
        turn.start.board.Tiles.Add(tile4);
        turn.start.board.Tiles.Add(tile5);
        turn.start.board.Tiles.Add(tile6);
        turn.start.board.Tiles.Add(tile7);
        turn.start.board.Tiles.Add(tile8);
        turn.start.board.Tiles.Add(tile9);

        Tile tile10 = new Tile();
        Tile tile11 = new Tile();
        Tile tile12 = new Tile();
        Tile tile13 = new Tile();
        Tile tile14 = new Tile();
        Tile tile15 = new Tile();
        Tile tile16 = new Tile();
        Tile tile17 = new Tile();
        Tile tile18 = new Tile();

        tile10.id = 10;
        tile10.tileType = "Water";
        tile10.x = 0;
        tile10.y = 0;
        tile10.z = 0;

        tile11.id = 11;
        tile11.tileType = "Fire";
        tile11.x = 0;
        tile11.y = 0;
        tile11.z = 1;

        tile12.id = 12;
        tile12.tileType = "Mountain";
        tile12.x = 0;
        tile12.y = 0;
        tile12.z = 2;

        tile13.id = 13;
        tile13.tileType = "Grass";
        tile13.x = 1;
        tile13.y = 0;
        tile13.z = 0;

        tile14.id = 14;
        tile14.tileType = "Water";
        tile14.x = 1;
        tile14.y = 0;
        tile14.z = 1;

        tile15.id = 15;
        tile15.tileType = "Dirt";
        tile15.x = 1;
        tile15.y = 0;
        tile15.z = 2;

        tile16.id = 16;
        tile16.tileType = "Fire";
        tile16.x = 2;
        tile16.y = 0;
        tile16.z = 0;

        tile17.id = 17;
        tile17.tileType = "Dirt";
        tile17.x = 2;
        tile17.y = 0;
        tile17.z = 1;

        tile18.id = 18;
        tile18.tileType = "Fire";
        tile18.x = 2;
        tile18.y = 0;
        tile18.z = 2;

        turn.end.board.Tiles.Add(tile10);
        turn.end.board.Tiles.Add(tile11);
        turn.end.board.Tiles.Add(tile12);
        turn.end.board.Tiles.Add(tile13);
        turn.end.board.Tiles.Add(tile14);
        turn.end.board.Tiles.Add(tile15);
        turn.end.board.Tiles.Add(tile16);
        turn.end.board.Tiles.Add(tile17);
        turn.end.board.Tiles.Add(tile18);

        clientInit(turn);
        //Todo: Connect to WebSocket
    }

    // Update is called once per frame
    void Update () {

        //Todo: Retrieve ProcessedTurn and pass it to processClient
        
        if (Input.GetKeyDown(KeyCode.I))
        {
            Console.Write("NO BUTTONS!");
            clientRefresh(turn);
        }
    }

    public void clientInit(ProcessedTurn Turn)
    {
        
            Board = Turn.start.board;
            Board.init();
        //foreach (Step step in Turn.steps)
        //{
        //    foreach (IAction action in step.actions)
        //    {
        //        switch (action.playerId)
        //        {
        //            case "Player1":
        //                processAction(Player1, action);
        //                break;
        //            case "Player2":
        //                processAction(Player2, action);
        //                break;
        //            case "Player3":
        //                processAction(Player3, action);
        //                break;
        //            case "Player4":
        //                processAction(Player4, action);
        //                break;
        //        }
        //    }
        //    // force render actions for this step
        //}

        //foreach(Card card in Turn.inputs)
        //{

        //}
        List<GameObject> inGameTiles = Board.Gametiles;
        Board = Turn.end.board;
        Board.Gametiles = inGameTiles;


    }

    public void clientRefresh(ProcessedTurn Turn)
    {
        Board.refresh();

    }


    private void processAction(Player player4, IAction action)
    {
        throw new NotImplementedException();
    }
}
