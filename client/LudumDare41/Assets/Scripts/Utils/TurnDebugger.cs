using System;
using System.Collections.Generic;
using Model;

namespace Utils {
    public static class TurnDebugger {
        private static Random rnd = new Random();
        private static string[] tileTypes = new string[]{"empty", "wall"};
        private static string[] cardTypes = new string[]{"MoveForward1", "MoveBackward1"};
        private static int boardWidth = 3;
        private static int boardHeight = 3;
            
        
        public static ProcessedTurn GenerateTurn() {
            var t = new ProcessedTurn();
            t.start = GenerateGameState();
            t.end = GenerateGameState();
            t.end.tick = t.start.tick + 1;
            // TODO: need to test steps
            t.steps = new List<Step>();
            // TODO: need to test inputs
            t.inputs = new List<Card>();
            return t;
        }

        public static GameState GenerateGameState() {
            var state = new GameState();
            state.tick = rnd.Next(10);
            state.gameBoard = GenerateGameBoard();
            state.players = new List<Player>();
            for (int i = 0; i < 4; i++) state.players.Add(GeneratePlayer());
            return state;
        }
        
        public static GameBoard GenerateGameBoard() {
            var b = new GameBoard();
            b.width = boardWidth;
            b.height = boardHeight;
            b.tiles = new List<Tile>();
            for (int x = 0; x < b.width; x++) {
                for (int y = 0; y < b.height; y++) {
                    var t = new Tile();
                    t.id = rnd.Next(100000);
                    t.pos = new Vector(x, y);
                    t.tileType = tileTypes[rnd.Next(tileTypes.Length)];
                    b.tiles.Add(t);
                }
            }
            return b;
        }

        public static Player GeneratePlayer() {
            var p = new Player();
            p.id = rnd.Next(10);
            p.name = "" + rnd.Next(10000);
            p.pos = new Vector();
            p.pos.x = rnd.Next(boardWidth);
            p.pos.y = rnd.Next(boardHeight);
            p.facing = new Vector();
            p.facing.x = rnd.Next(3) - 1;
            p.facing.y = rnd.Next(3) - 1;
            p.hand = new List<Card>();
            for (int i = 0; i < 5; i++) p.hand.Add(GenerateCard());
            return p;
        }

        public static Card GenerateCard() {
            var c = new Card();
            c.id = rnd.Next(100000);
            c.priority = rnd.Next(1000000);
            c.cardType = cardTypes[rnd.Next(cardTypes.Length)];
            return c;
        }
    }
}