package main

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	numOfPlayers = 4
	numOfSquares = 24
)

type pos struct {
	i, j int
}

type GameBoard struct {
	squares []int
}

type Player struct {
	Name     string
	Position int
	Color    string
	Chess    int
}

var (
	players = []Player{
		{Name: "Player 1", Position: 0, Color: "Red", Chess: 4},
		{Name: "Player 2", Position: 0, Color: "Green", Chess: 4},
		{Name: "Player 3", Position: 0, Color: "Blue", Chess: 4},
		{Name: "Player 4", Position: 0, Color: "Yellow", Chess: 4},
	}
)

func rollDice() int {
	var roll_1 int = rand.Intn(6) + 1
	var roll_2 int = rand.Intn(6) + 1
	if rand.Intn(2) == 0 {
		fmt.Printf("You rolled (+): %d\t", roll_1+roll_2)
		return roll_1 + roll_2
	} else {
		if roll_1-roll_2 > 0 {
			fmt.Printf("You rolled (-): %d\t", roll_1-roll_2)
			return roll_1 - roll_2
		} else {
			fmt.Printf("You rolled (-): %d\t", roll_1-roll_2)
			return 0
		}
	}
}

func initGameBoard(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	GameBoard := make([][]int, row)
	for i := range GameBoard {
		GameBoard[i] = make([]int, col)
		for j := range GameBoard[i] {
			fmt.Fscanf(file, "%d", &GameBoard[i][j])
		}
	}
	return GameBoard
}

func movePlayer(player *Player, diceRoll int) {
	player.Position = player.Position + diceRoll
}

var dirs = [4]pos{
	{-1, 0}, // abajo
	{1, 0},  // arriba
	{0, -1}, // iquierda
	{0, 1},  // derecha
}

func (p pos) move(r pos) pos {
	return pos{p.i + r.i, p.j + r.j}
}

func move(GameBoard [][]int, player1, player2, player3, player4 chan pos) {

}

func main() {
	var currentPlayerIndex int

	GameBoard := initGameBoard("GameBoard.in")

	for i := range GameBoard {
		for j := range GameBoard[i] {
			fmt.Printf("%d\t", GameBoard[i][j])
		}
		fmt.Println()
	}
	// end := pos{int(len(GameBoard) / 2), int(len(GameBoard) / 2)}
	Players := players
	for {
		currentPlayer := &Players[currentPlayerIndex]
		fmt.Printf("It's %s's turn.\n", currentPlayer.Name)

		movePlayer(currentPlayer, rollDice())
		fmt.Printf("You moved to square %d.\n", currentPlayer.Position)

		if currentPlayer.Position >= numOfSquares {
			currentPlayer.Position = 0
			fmt.Printf("%s: Won 1 Race.\n", currentPlayer.Name)
			currentPlayer.Chess--
			fmt.Printf("%s: Remain Chess %d.\n", currentPlayer.Name, currentPlayer.Chess)
			if currentPlayer.Chess == 0 {
				fmt.Printf("%s --- Won Game ---\n", currentPlayer.Name)
				break
			}
		}

		currentPlayerIndex = (currentPlayerIndex + 1) % numOfPlayers
	}
}
