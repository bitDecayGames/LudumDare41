package game

import (
	"fmt"
	"log"
	"math/rand"
	"sort"

	"github.com/bitDecayGames/LudumDare41/server/utils"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/logic"
	"github.com/bitDecayGames/LudumDare41/server/state"
)

const (
	HAND_SIZE     = 5
	PLAYER_COLORS = 4
)

var NextPlayerColor = 0

func GetNextPlayerTeam() int {
	next := NextPlayerColor
	NextPlayerColor = (NextPlayerColor + 1) % PLAYER_COLORS
	return next
}

type Game struct {
	Name string
	// WARNING This is only initial state, do not read
	Players map[string]*state.Player
	Board   gameboard.GameBoard
	CardSet cards.CardSet

	CurrentState state.GameState
	LastSequence logic.StepSequence

	PreviousState state.GameState

	pendingSubmissions map[string][]cards.Card // Player submissions
	pendingSequence    []cards.Card            // Ordered list of all player cards
}

func newGame(players map[string]*state.Player, board gameboard.GameBoard, cardSet cards.CardSet, name string) *Game {

	playerNum := 1
	for _, player := range players {
		player.Deck = cards.NewDeckFromSet(cardSet, len(players), playerNum)
		player.Team = GetNextPlayerTeam()
		playerNum += 1
	}

	currentState := state.NewState(-1, players, board)
	currentState.Crate = utils.DeadVector
	currentState.NextCrate = utils.DeadVector

	fmt.Println(fmt.Sprintf("New State: %+v", currentState))

	return &Game{
		Name:               name,
		Players:            players,
		Board:              board,
		CardSet:            cardSet,
		PreviousState:      currentState,
		CurrentState:       DealCards(currentState),
		pendingSubmissions: make(map[string][]cards.Card),
	}
}

func DealCards(inState state.GameState) state.GameState {
	priority := 1
	for i, _ := range inState.Players {
		for len(inState.Players[i].Hand) < HAND_SIZE {
			if len(inState.Players[i].Deck) == 0 {
				inState.Players[i].Deck = cards.ShuffleCards(inState.Players[i].Discard)
				inState.Players[i].Discard = make([]cards.Card, 0)
			}
			drawnCard := inState.Players[i].Deck[0]
			drawnCard.Owner = inState.Players[i].Name
			inState.Players[i].Deck = inState.Players[i].Deck[1:]
			inState.Players[i].Hand = append(inState.Players[i].Hand, drawnCard)
			priority += 1
		}
	}
	return inState
}

func (g *Game) AreSubmissionsComplete() bool {
	numSubmissons := len(g.pendingSubmissions)
	log.Printf("%v/%v player submissions are pending", numSubmissons, len(g.Players))
	return numSubmissons == len(g.Players)
}

func (g *Game) GetPlayer(name string) (state.Player, error) {
	for _, p := range g.CurrentState.Players {
		if p.Name == name {
			return p, nil
		}
	}

	return state.Player{}, fmt.Errorf("player not fround with name %s", name)
}

func (g *Game) SubmitCards(playerName string, tick int, cardIds []int) error {
	log.Printf("Player %s is submiiting card IDs %+v for tick %v", playerName, cardIds, tick)

	if g.CurrentState.Tick != tick {
		return fmt.Errorf("expected tick of %v, not %v", g.CurrentState.Tick, tick)
	}

	// TODO validate these cards
	if g.pendingSubmissions[playerName] != nil {
		return fmt.Errorf("Player already has a pending submission")
	}

	player, err := g.GetPlayer(playerName)
	if err != nil {
		return err
	}

	log.Printf("Player %s's current hand: %+v", playerName, player.Hand)

	// Find cards
	submission := []cards.Card{}
	for _, id := range cardIds {
		for _, card := range player.Hand {
			if card.ID == id {
				log.Printf("Found matching card for id %v: %+v", id, card)
				submission = append(submission, card)
			}
		}
	}

	g.pendingSubmissions[playerName] = submission
	log.Printf("Player %s submitted cards: %+v", playerName, submission)

	return nil
}

func (g *Game) AggregateTurn() []cards.Card {
	g.pendingSequence = make([]cards.Card, 0)
	cardsAdded := true
	for cardsAdded {
		cardsAdded = false
		cardOrder := make([]cards.Card, 0)
		for name, pendingCards := range g.pendingSubmissions {
			if len(pendingCards) == 0 {
				delete(g.pendingSubmissions, name)
				continue
			}
			cardOrder = append(cardOrder, pendingCards[0])
			g.pendingSubmissions[name] = pendingCards[1:]
			cardsAdded = true
		}
		sort.Slice(cardOrder, func(i, j int) bool { return cardOrder[i].Priority > cardOrder[j].Priority })
		g.pendingSequence = append(g.pendingSequence, cardOrder...)
	}
	return g.pendingSequence
}

func (g *Game) ExecuteTurn() {
	// This should carry out the full step sequence (cards) and calculate all actions that fall out

	// 0. Set previous state
	g.PreviousState = g.CurrentState
	// 1. Get starting state
	startState := getCopy(g.CurrentState)
	// 2. Execute all cards
	intermState := g.CurrentState
	stepSequence := logic.StepSequence{
		Cards: g.pendingSequence,
	}
	for _, c := range g.pendingSequence {
		fmt.Println(fmt.Sprintf("%+v", c))
		newSteps, newState := logic.ApplyCard(c, intermState)
		stepSequence.Steps = append(stepSequence.Steps, newSteps...)

		intermState = newState
	}

	// respawn any dead players.  This assumes zero downtime -- you die, you respawn instantly
	steps, intermState := respawnObjects(intermState)
	stepSequence.Steps = append(stepSequence.Steps, steps...)

	intermState = DealCards(intermState)
	// 3. Update clients with these things:
	fmt.Println(startState)
	fmt.Println(stepSequence)
	fmt.Println(fmt.Sprintf("Pending Seq %+v", g.pendingSequence))
	intermState.Tick += 1
	g.LastSequence = stepSequence
	g.CurrentState = intermState
}

func getCopy(g state.GameState) state.GameState {
	newState := state.GameState{
		Tick:      g.Tick,
		Players:   make([]state.Player, len(g.Players)),
		Crate:     g.Crate,
		NextCrate: g.NextCrate,
		Board:     g.Board,
	}

	for i, _ := range g.Players {
		newState.Players[i] = g.Players[i]
	}

	return newState
}

func respawnObjects(g state.GameState) ([]logic.Step, state.GameState) {
	steps := make([]logic.Step, 0)

	// Get empty tiles
	emptyTiles := g.Board.GetTilesByType(gameboard.Empty_tile)
	// Randomly sort them
	rand.Shuffle(len(emptyTiles), func(i, j int) {
		emptyTiles[i], emptyTiles[j] = emptyTiles[j], emptyTiles[i]
	})

	// Flag tiles other players are on
	for i, et := range emptyTiles {
		for _, p := range g.Players {
			if utils.VecEquals(et.Pos, p.Pos) {
				et.TempOccupied = true
				emptyTiles[i] = et
			}
		}
	}

	// do crate logic before we respawn players since crates coming in might kill players
	crateSteps, g, emptyTiles := manageCrates(g, emptyTiles)
	if len(crateSteps) > 0 {
		steps = append(steps, crateSteps...)
	}

	spawnPlayerStep, g, emptyTiles := spawnDeadPlayers(g, emptyTiles)
	if len(spawnPlayerStep.Actions) > 0 {
		steps = append(steps, spawnPlayerStep)
	}

	return steps, g
}

func spawnDeadPlayers(g state.GameState, emptyTiles []gameboard.Tile) (logic.Step, state.GameState, []gameboard.Tile) {
	step := logic.Step{}
	for i, p := range g.Players {
		if utils.VecEquals(p.Pos, utils.DeadVector) {
			// Find tiles to place respawned players on.
			// NOTE This could hit an unsolvable situation
			for k, tile := range emptyTiles {
				if !tile.TempOccupied {
					log.Printf("Respawning player %s at %+v", p.Name, tile.Pos)
					g.Players[i].Pos = tile.Pos
					step.Actions = append(step.Actions, logic.GetAction(logic.Action_spawn, p.Name, tile.Pos))

					tile.TempOccupied = true
					emptyTiles[k] = tile

					break
				}
			}
		}
	}
	return step, g, emptyTiles
}

func manageCrates(g state.GameState, emptyTiles []gameboard.Tile) ([]logic.Step, state.GameState, []gameboard.Tile) {
	steps := make([]logic.Step, 0)
	if utils.VecEquals(g.NextCrate, utils.DeadVector) {
		// first time through. Set this and wait one turn before spawning actual crate
		for k, tile := range emptyTiles {
			if !tile.TempOccupied {
				log.Printf("Initializing next crate position to %+v", tile.Pos)
				g.Crate = tile.Pos
				steps = append(steps, logic.Step{
					Actions: []logic.Action{logic.GetAction(logic.Action_set_next_crate, "gameBoard", tile.Pos)},
				})
				tile.TempOccupied = true
				emptyTiles[k] = tile
				return steps, g, emptyTiles
			}
		}
		panic("No place to put the next crate")
	}

	if utils.VecEquals(g.Crate, utils.DeadVector) {

		crateBoomStep := logic.Step{}
		playerKillStep := logic.Step{}
		for x := g.NextCrate.X - 1; x <= g.NextCrate.X+1; x++ {
			for y := g.NextCrate.Y - 1; y <= g.NextCrate.Y+1; y++ {
				tileLoc := utils.NewVec(x, y)
				if g.Board.OnBoard(tileLoc) {
					crateBoomStep.Actions = append(crateBoomStep.Actions, logic.GetAction(logic.Action_tile_explode, "gameBoard", tileLoc))

					if logic.IsPlayerOccupying(tileLoc, g) {
						killP := logic.GetPlayerAtPos(tileLoc, g)
						playerKillStep.Actions = append(playerKillStep.Actions, logic.GetAction(logic.Action_death, killP.Name, killP.Pos))
					}
				}
			}
		}

		if len(crateBoomStep.Actions) > 0 {
			steps = append(steps, crateBoomStep)
		}

		if len(playerKillStep.Actions) > 0 {
			steps = append(steps, playerKillStep)
		}

		g.Crate = g.NextCrate
		steps = append(steps, logic.Step{
			Actions: []logic.Action{logic.GetAction(logic.Action_drop_crate, "gameBoard", g.Crate)},
		})

		for k, tile := range emptyTiles {
			if !tile.TempOccupied {
				log.Printf("Setting next crate pos to %+v", tile.Pos)
				g.NextCrate = tile.Pos
				steps = append(steps, logic.Step{
					Actions: []logic.Action{logic.GetAction(logic.Action_set_next_crate, "gameBoard", tile.Pos)},
				})
				tile.TempOccupied = true
				emptyTiles[k] = tile
				break
			}
		}
	}

	return steps, g, emptyTiles
}
