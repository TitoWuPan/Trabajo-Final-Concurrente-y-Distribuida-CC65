package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

type Direction int

const (
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
	//id del jugador
	Name string
	//posicion del jugador en el tablero
	Position pos
	//Color de la pieza
	Color string

	Chess int
	//
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

// agregar concurrencia goroutine independinete de tirar dado y que est sea el agrupador
func rollDice() int {
	var roll_1 int = rand.Intn(6) + 1 //roll_1= go roll_dice()
	var roll_2 int = rand.Intn(6) + 1 //roll_2=go roll_dice()
	if rand.Intn(2) == 0 {
		fmt.Printf("You rolled (+): %d\n", roll_1+roll_2)
		return roll_1 + roll_2
	} else {
		if roll_1-roll_2 > 0 {
			fmt.Printf("You rolled (-): %d\n", roll_1-roll_2)
			return roll_1 - roll_2
		} else {
			fmt.Printf("You rolled (-): %d\n", roll_1-roll_2)
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

func exitCheck(curPos pos) bool {
	if curPos.i == GameBoard.rows-2 && curPos.j == GameBoard.columns-2 {
		return true
	} else {
		return false
	}
}

func move(players Player, dice int, dir Direction) Player {
	players.Direction = int(dir)

	fmt.Printf("%s at (%d, %d) in direction %d\n", players.Name, players.Position.i, players.Position.j, players.Direction)

	for i := 0; i < dice; i++ {
		if Direction(players.Direction) == Up {
			//look left
			if GameBoard.maze[players.Position.i][players.Position.j-1] != 1 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
				//look up
			} else if GameBoard.maze[players.Position.i-1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
				//look right
			} else if GameBoard.maze[players.Position.i][players.Position.j+1] != 1 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
				//look down
			} else if GameBoard.maze[players.Position.i+1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
			}
		} else if Direction(players.Direction) == Left {
			//look down
			if GameBoard.maze[players.Position.i+1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
				//look left
			} else if GameBoard.maze[players.Position.i][players.Position.j-1] != 1 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
				//look up
			} else if GameBoard.maze[players.Position.i-1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
				//look right
			} else if GameBoard.maze[players.Position.i][players.Position.j+1] != 1 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
			}
		} else if Direction(players.Direction) == Right {
			//look up
			if GameBoard.maze[players.Position.i-1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
				//look right
			} else if GameBoard.maze[players.Position.i][players.Position.j+1] != 1 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
				//look down
			} else if GameBoard.maze[players.Position.i+1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
				// look left
			} else if GameBoard.maze[players.Position.i][players.Position.j-1] != 1 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
			}
		} else if Direction(players.Direction) == Down {
			//look right
			if GameBoard.maze[players.Position.i][players.Position.j+1] != 1 {
				players.Position = players.Position.move(direction[0].Right)
				players.Direction = int(Right)
				//look down
			} else if GameBoard.maze[players.Position.i+1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Down)
				players.Direction = int(Down)
				//look left
			} else if GameBoard.maze[players.Position.i][players.Position.j-1] != 1 {
				players.Position = players.Position.move(direction[0].Left)
				players.Direction = int(Left)
				//look up
			} else if GameBoard.maze[players.Position.i-1][players.Position.j] != 1 {
				players.Position = players.Position.move(direction[0].Up)
				players.Direction = int(Up)
			}
		}

		if GameBoard.maze[players.Position.i][players.Position.j] == 2 {
			fmt.Printf("%s Moving at (%d, %d) in direction %d\t", players.Name, players.Position.i, players.Position.j, players.Direction)
			fmt.Printf("%s Fall in Tramp.\n", players.Name)
			return players
		}

		if exitCheck(players.Position) {
			fmt.Printf("%s Moving at (%d, %d) in direction %d\n", players.Name, players.Position.i, players.Position.j, players.Direction)
			return players
		}
	}
	fmt.Printf("%s Moving at (%d, %d) in direction %d\n", players.Name, players.Position.i, players.Position.j, players.Direction)
	return players
}

func play(player1, player2, player3, player4 Player) {
	var wg sync.WaitGroup

	p1 := player1
	p2 := player2
	p3 := player3
	p4 := player4

	for {
		wg.Add(4)

		var p1_new Player = move(p1, rollDice(), Direction(p1.Direction))
		wg.Done()

		var p2_new Player = move(p2, rollDice(), Direction(p2.Direction))
		wg.Done()

		var p3_new Player = move(p3, rollDice(), Direction(p3.Direction))
		wg.Done()

		var p4_new Player = move(p4, rollDice(), Direction(p4.Direction))
		wg.Done()

		p1 = p1_new
		p2 = p2_new
		p3 = p3_new
		p4 = p4_new

		if exitCheck(p1.Position) {
			p1.Chess--
			p1.Position = pos{GameBoard.startRow, GameBoard.startColumn}
			fmt.Printf("%s Finish 1 run, Rest (%d) chess. \n", p1.Name, p1.Chess)
			if p1.Chess == 0 {
				fmt.Printf("%s Win\n", p1.Name)
				break
			}
		}

		if exitCheck(p2.Position) {
			p2.Chess--
			p2.Position = pos{GameBoard.startRow, GameBoard.startColumn}
			fmt.Printf("%s Finish 1 run, Rest (%d) chess. \n", p2.Name, p2.Chess)
			if p2.Chess == 0 {
				fmt.Printf("%s Win\n", p2.Name)
				break
			}
		}

		if exitCheck(p3.Position) {
			p3.Chess--
			p3.Position = pos{GameBoard.startRow, GameBoard.startColumn}
			fmt.Printf("%s Finish 1 run, Rest (%d) chess. \n", p3.Name, p3.Chess)
			if p3.Chess == 0 {
				fmt.Printf("%s Win\n", p3.Name)
				break
			}
		}

		if exitCheck(p4.Position) {
			p4.Chess--
			p4.Position = pos{GameBoard.startRow, GameBoard.startColumn}
			fmt.Printf("%s Finish 1 run, Rest (%d) chess. \n", p4.Name, p4.Chess)
			if p4.Chess == 0 {
				fmt.Printf("%s Win\n", p4.Name)
				break
			}
		}
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

	play(players[0], players[1], players[2], players[3])
}
