package main

import (
	"fmt"
	"hunt-the-wumpus/util"
	"math/rand"
)

type GameStatus int

type GameState struct {
	flags                   int
	x, y                    int
	wHealth                 int
	wX, wY                  int
	arrows                  int
	pits, arrow_cells, bats [][2]int
	status                  GameStatus
}

func NewGameState(flags util.Flags) GameState {
	var state GameState
	state.x, state.y = randomCell(flags)
	state.wHealth = *flags.WHP
	state.wX, state.wY = randomCell(flags)
	state.arrows = *flags.Arrows
	state.pits, state.arrow_cells, state.bats = generateBotomlessPitsBatsAndArrows(state.x, state.y, flags)
	return state
}

// Dy default each cell has a:
// - 10% change to have a pit
// - 10% chance to have an arrow
// - 5% chance to have bats
func generateBotomlessPitsBatsAndArrows(x int, y int, flags util.Flags) ([][2]int, [][2]int, [][2]int) {
	pits := [][2]int{}
	arrows := [][2]int{}
	bats := [][2]int{}

	for i := 0; i < *flags.Rows; i++ {
		for j := 0; j < *flags.Cols; j++ {
			if x == i && y == j {
				continue
			}

			roll := rand.Int() % 100
			if roll < *flags.BpChance {
				pit := [2]int{i, j}
				pits = append(pits, pit)
			}

			if roll >= *flags.BatsChance && roll < *flags.BpChance+*flags.ArrowChance {
				arrow := [2]int{i, j}
				arrows = append(arrows, arrow)
			}

			if roll >= *flags.BpChance+*flags.ArrowChance && roll < *flags.BpChance+*flags.ArrowChance+*flags.BatsChance {
				bat := [2]int{i, j}
				bats = append(bats, bat)
			}
		}
	}

	return pits, arrows, bats
}

func randomCell(flags util.Flags) (int, int) {
	return rand.Int() % *flags.Rows, rand.Int() % *flags.Cols
}

const (
	InProgress GameStatus = iota
	Won
	Lost
)

func GetGameStatus(state GameState) GameStatus {
	if state.wHealth == 0 {
		fmt.Println("You win!")
		return Won
	}

	if state.arrows == 0 {
		fmt.Print("No arrows left.")
		return Lost
	}

	if state.x == state.wX && state.y == state.wY {
		fmt.Println("Wumpus ate you.")
		return Lost
	}

	for _, pit := range state.pits {
		if state.x == pit[0] && state.y == pit[1] {
			fmt.Println("You fell in a bottomless pit.")
			return Lost
		}
	}

	return InProgress
}
