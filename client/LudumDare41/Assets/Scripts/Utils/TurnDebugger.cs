using System;
using System.Collections.Generic;
using Model;

namespace Utils {
    public static class TurnDebugger {
        private static Random rnd = new Random();
        private static int boardWidth = 5;
        private static int boardHeight = 5;
            
        
        public static ProcessedTurn GenerateTurn() {
            var t = new ProcessedTurn();
            t.start = GenerateGameState();
            t.end = GenerateGameState();
            t.end.tick = t.start.tick + 1;
            t.diff = GenerateStepSequence(t.start.players[0].name, "spawnAction", "moveEastAction", "moveSouthAction");
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
                    t.tileType = Tile.TILE_TYPES[rnd.Next(Tile.TILE_TYPES.Length)];
                    b.tiles.Add(t);
                }
            }
            return b;
        }

        public static Player GeneratePlayer() {
            var p = new Player();
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
            c.cardType = Card.CARD_TYPES[rnd.Next(Card.CARD_TYPES.Length)];
            return c;
        }

        public static StepSequence GenerateStepSequence(string playerName, params string[] actions) {
            var seq = new StepSequence();
            seq.cards = new List<Card>();
            for (int i = 0; i < 5; i++) seq.cards.Add(GenerateCard());
            seq.steps = new List<Step>();
            for (int i = 0; i + 1 < actions.Length; i += 2) {
                var first = actions[i];
                var second = actions[i + 1];
                if (i % 4 == 0 && i + 2 < actions.Length) seq.steps.Add(GenerateStep(playerName, first, second));
                else {
                    seq.steps.Add(GenerateStep(playerName, first));
                    seq.steps.Add(GenerateStep(playerName, second));
                }
            }
            if (actions.Length % 2 != 0) seq.steps.Add(GenerateStep(playerName, actions[actions.Length - 1]));
            return seq;
        }

        public static Step GenerateStep(string playerName, params string[] actions) {
            var step = new Step();
            step.actions = new List<ActionData>();
            foreach (string act in actions) {
                step.actions.Add(GenerateActionData(playerName, act));
            }
            return step;
        }

        public static ActionData GenerateActionData(string playerName, string actionType) {
            var action = new ActionData();
            action.id = rnd.Next(10);
            action.actionType = actionType;
            action.playerId = playerName;
            return action;
        }
        
    }
}