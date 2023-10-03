package main

import (
	"fmt"
	"hunt-the-wumpus/util"
	"math/rand"
)

// left, right, up, down
var simpleMove = [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func move(x int, y int, dir int, flags util.Flags) (int, int) {
	newX := x + simpleMove[dir][0]
	newY := y + simpleMove[dir][1]

	if x < 0 || x >= *flags.Rows {
		newX = x
	}

	if y < 0 || y >= *flags.Cols {
		newY = y
	}

	return newX, newY
}

func moveWumpus(x int, y int, flags util.Flags) (int, int) {
	move := rand.Int() % 4
	newX := x + simpleMove[move][0]
	newY := y + simpleMove[move][1]

	if newX < 0 || newX >= *flags.Rows {
		newX = x
	}

	if newY < 0 || newY >= *flags.Cols {
		newY = y
	}

	return newX, newY
}

var arrowSymbol = [4]string{"<<", ">>", "/\\", "\\/"}

func shoot(x int, y int, dir int, wX int, wY int, arrows int, flags util.Flags) bool {
	i := x + simpleMove[dir][0]
	j := y + simpleMove[dir][1]

	for i < *flags.Rows && j < *flags.Cols && i >= 0 && j >= 0 {
		if i == wX && j == wY {
			util.PrintBoard([][2]int{{x, y}, {i, j}}, []string{"🏹", "🐻"}, arrows, flags)
			return true
		}

		util.PrintBoard([][2]int{{i, j}, {x, y}}, []string{arrowSymbol[dir], "🏹"}, arrows, flags)

		i = i + simpleMove[dir][0]
		j = j + simpleMove[dir][1]
	}

	return false
}

func feelingSomething(x int, y int, cells [][2]int) bool {
	for _, cell := range cells {
		for _, add := range simpleMove {
			if x+add[0] == cell[0] && y+add[1] == cell[1] {
				return true
			}
		}
	}
	return false
}

func checkForFeelings(gs GameState) {
	if feelingSomething(gs.x, gs.y, gs.pits) {
		fmt.Println("You feel a breeze.")
	}

	if feelingSomething(gs.x, gs.y, gs.bats) {
		fmt.Println("You hear flapping.")
	}

	if feelingSomething(gs.x, gs.y, [][2]int{{gs.wX, gs.wY}}) {
		fmt.Println("You smell a wumpus.")
	}
}

func somethingInCell(x int, y int, cells [][2]int) bool {
	for _, cell := range cells {
		if x == cell[0] && y == cell[1] {
			return true
		}
	}
	return false
}

const (
	GoCommand   string = "go"
	FireCommand string = "fire"
)

func handleCommand(gs GameState, flags util.Flags) GameState {
	fmt.Print("Your move: ")
	command, dir := util.ReadCommandAndDirrection()

	if command == GoCommand {
		gs.x, gs.y = move(gs.x, gs.y, dir, flags)

		if somethingInCell(gs.x, gs.y, gs.arrow_cells) {
			fmt.Println("Arrow found!")
			gs.arrows++
		}

		if somethingInCell(gs.x, gs.y, gs.bats) {
			fmt.Println("Bats move you!")
			gs.x, gs.y = randomCell(flags)
		}
	}

	if command == FireCommand {
		if shoot(gs.x, gs.y, dir, gs.wX, gs.wY, gs.arrows, flags) {
			gs.wHealth--
		}
		gs.arrows--
	}

	gs.status = GetGameStatus(gs)
	return gs
}

func main() {
	flags := util.ParseFlags()

	gameState := NewGameState(flags)

	for gameState.status == InProgress {
		util.PrintBoard([][2]int{{gameState.x, gameState.y}}, []string{"🏹"}, gameState.arrows, flags)

		checkForFeelings(gameState)

		gameState = handleCommand(gameState, flags)

		gameState.wX, gameState.wY = moveWumpus(gameState.wX, gameState.wY, flags)
	}

	if gameState.status == Won {
		fmt.Println("You won!")
	}

	if gameState.status == Lost {
		fmt.Println("You Lost!")
	}
}
