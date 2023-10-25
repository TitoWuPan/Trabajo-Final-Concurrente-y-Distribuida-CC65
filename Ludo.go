package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Direction int

const (
	numOfPlayers = 4
	numOfSquares = 24

	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

type pos struct {
	i, j int
}

var GameBoard struct {
	columns     int
	rows        int
	startRow    int
	startColumn int
	maze        [][]int
}

type Player struct {
	Name      string
	Position  pos
	Color     string
	Chess     int
	Direction int
}

type dir struct {
	Up    pos
	Down  pos
	Left  pos
	Right pos
}

var (
	players = []Player{
		{Name: "Player 1", Position: pos{1, 1}, Color: "Red", Chess: 4, Direction: int(Up)},
		{Name: "Player 2", Position: pos{1, 1}, Color: "Green", Chess: 4, Direction: int(Up)},
		{Name: "Player 3", Position: pos{1, 1}, Color: "Blue", Chess: 4, Direction: int(Up)},
		{Name: "Player 4", Position: pos{1, 1}, Color: "Yellow", Chess: 4, Direction: int(Up)},
	}
	direction = []dir{
		{Up: pos{-1, 0}, // Arriba
			Down:  pos{1, 0},  // Abajo
			Left:  pos{0, -1}, // Izquierda
			Right: pos{0, 1},  // Derecha
		},
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

func initGameBoard(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fmt.Fscanf(file, "%d %d %d %d", &GameBoard.rows, &GameBoard.columns, &GameBoard.startRow, &GameBoard.startColumn)

	GameBoard.maze = make([][]int, GameBoard.rows)
	for i := 0; i < GameBoard.rows; i++ {
		GameBoard.maze[i] = make([]int, GameBoard.columns)
		for j := 0; j < GameBoard.columns; j++ {
			fmt.Fscanf(file, "%d", &GameBoard.maze[i][j])
		}
	}
}

func (p pos) move(r pos) pos {
	return pos{p.i + r.i, p.j + r.j}
}

func exitCheck(curPos pos, playerNum int) bool {
	if curPos.i == GameBoard.rows-2 && curPos.j == GameBoard.columns-2 {
		players[playerNum].Chess--
		fmt.Printf("%s Found the Exit (%d, %d), Rest %d Chess to Finish\n", players[playerNum].Name, players[playerNum].Position.i, players[playerNum].Position.j, players[playerNum].Chess)
		players[playerNum].Position = pos{GameBoard.startRow, GameBoard.startColumn}
		return true
	} else {
		return false
	}
}

func Check(curPos pos, playerNum int) {
	players[playerNum].Position = curPos
}

func move(players Player, dice int, playerNum int) {

	for i := 0; i < dice; i++ {
		if Direction(players.Direction) == Up {
			//look left
			if GameBoard.maze[players.Position.i][players.Position.j-1] == 0 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
				//look up
			} else if GameBoard.maze[players.Position.i-1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
				//look right
			} else if GameBoard.maze[players.Position.i][players.Position.j+1] == 0 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
				//look down
			} else if GameBoard.maze[players.Position.i+1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
			}
		} else if Direction(players.Direction) == Left {
			//look down
			if GameBoard.maze[players.Position.i+1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
				//look left
			} else if GameBoard.maze[players.Position.i][players.Position.j-1] == 0 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
				//look up
			} else if GameBoard.maze[players.Position.i-1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
				//look right
			} else if GameBoard.maze[players.Position.i][players.Position.j+1] == 0 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
			}
		} else if Direction(players.Direction) == Right {
			//look right
			if GameBoard.maze[players.Position.i][players.Position.j+1] == 0 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
				//look up
			} else if GameBoard.maze[players.Position.i-1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
				//look down
			} else if GameBoard.maze[players.Position.i+1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
				// look left
			} else if GameBoard.maze[players.Position.i][players.Position.j-1] == 0 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
			}
		} else if Direction(players.Direction) == Down {
			//look right
			if GameBoard.maze[players.Position.i][players.Position.j+1] == 0 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
				//look down
			} else if GameBoard.maze[players.Position.i+1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
				//look left
			} else if GameBoard.maze[players.Position.i][players.Position.j-1] == 0 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
				//look up
			} else if GameBoard.maze[players.Position.i-1][players.Position.j] == 0 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
			}
		}
		// fmt.Printf("%s Moving at (%d, %d)\n", players.Name, players.Position.i, players.Position.j)
		// fmt.Printf("%d\n", Direction(players.Direction))
		// Check(players.Position, playerNum)
		// if exitCheck(players.Position, playerNum) { break }
	}
}

func play(player1, player2, player3, player4 chan Player) {
	p1 := <-player1
	p2 := <-player2
	p3 := <-player3
	p4 := <-player4

	for {

		player1 <- p1
		player2 <- p2
		player3 <- p3
		player4 <- p4

		close(player1)
		close(player2)
		close(player3)
		close(player4)
	}
}

func main() {
	initGameBoard("GameBoard.in")

	for i := range GameBoard.maze {
		for j := range GameBoard.maze[i] {
			fmt.Printf("%d\t", GameBoard.maze[i][j])
		}
		fmt.Println()
	}

	move(players[0], 100, 0)
	fmt.Printf("%s Win (%d, %d)\n", players[0].Name, players[0].Position.i, players[0].Position.j)
}
