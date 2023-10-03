package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// helper string to build the board
const rowSeperator = "----+"
const colSeperator = "    |"

// time between board renders
const timeToSleep = 400 * time.Millisecond

func PrintBoard(coords [][2]int, emotes []string, arrows int, flags Flags) {
	RefreshConsole()

	fmt.Println("Arrows left: ", arrows)

	for i := 0; i < *flags.Rows; i++ {
		topRow := "+"
		for j := 0; j < *flags.Cols; j++ {
			topRow += rowSeperator
		}
		fmt.Println(topRow)

		row := "|"
		for j := 0; j < *flags.Cols; j++ {
			row += colSeperator
		}

		for j := 0; j < len(emotes); j++ {
			if coords[j][0] == i {
				row = row[:(coords[j][1]*5+2)] + emotes[j] + row[(coords[j][1]*5+4):]
			}
		}

		fmt.Println(row)
	}

	topRow := "+"
	for j := 0; j < *flags.Cols; j++ {
		topRow += rowSeperator
	}
	fmt.Println(topRow)
}

func getDir(dirText string) int {
	var dir int
	switch dirText {
	case "left":
		dir = 0
	case "right":
		dir = 1
	case "up":
		dir = 2
	case "down":
		dir = 3
	}
	return dir
}

func ReadCommandAndDirrection() (string, int) {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.Trim(line, "\n")
	command := strings.Split(line, " ")
	return command[0], getDir(command[1])
}

func RefreshConsole() {
	time.Sleep(timeToSleep)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
