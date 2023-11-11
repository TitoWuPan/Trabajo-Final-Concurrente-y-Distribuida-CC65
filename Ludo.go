package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Direction int

var remotehost string
var playern int = 0

const (
	PlayerNum           = 10
	Up        Direction = 0
	Down      Direction = 1
	Left      Direction = 2
	Right     Direction = 3
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
	Pieces    int
	Direction int
	Turno     int
}

type dir struct {
	Up    pos
	Down  pos
	Left  pos
	Right pos
}

var (
	direction = []dir{{Up: pos{-1, 0}, Down: pos{1, 0}, Left: pos{0, -1}, Right: pos{0, 1}}}
	players   []*Player
	mu        sync.Mutex
)

func rollDices() int {
	mu.Lock()
	roll_1 := rand.Intn(6) + 1
	roll_2 := rand.Intn(6) + 1
	roll_3 := rand.Intn(2)
	mu.Unlock()

	if roll_3 == 0 {
		return roll_1 + roll_2
	} else {
		if roll_1-roll_2 > 0 {
			return roll_1 - roll_2
		} else {
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

	reader := bufio.NewReader(file)
	for i := 0; i < GameBoard.rows; i++ {
		GameBoard.maze[i] = make([]int, GameBoard.columns)
		fmt.Fscanf(reader, "%d")
		for j := 0; j < GameBoard.columns; j++ {
			fmt.Fscanf(reader, "%d", &GameBoard.maze[i][j])
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

func move(players *Player, dice int, dir Direction) *Player {
	players.Direction = int(dir)
	fmt.Printf("%s at (%d, %d)\n", players.Name, players.Position.i, players.Position.j)

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
			fmt.Printf("%s Moving at (%d, %d)\t ->   %s Fall in Tramp.\n", players.Name, players.Position.i, players.Position.j, players.Name)
			return players
		}

		if exitCheck(players.Position) {
			fmt.Printf("%s Moving at (%d, %d)\n", players.Name, players.Position.i, players.Position.j)
			return players
		}
	}
	fmt.Printf("%s Moving at (%d, %d)\n", players.Name, players.Position.i, players.Position.j)
	return players
}

var Playerhost string
var Destinyhost string

func Set_Players(players []*Player) {
	con, _ := net.Dial("tcp", Destinyhost)

	arrBytesJson, _ := json.Marshal(players)
	strMsgJson := string(arrBytesJson)

	fmt.Fprintln(con, strMsgJson)
}

func handle(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	str, _ := br.ReadString('\n')

	json.Unmarshal([]byte(str), &players)

	// fmt.Printf("Get %s\n", players[playern].Name)
	var player = move(players[playern], rollDices(), Direction(players[playern].Direction))

	if exitCheck(player.Position) {
		player.Pieces--
		player.Position = pos{GameBoard.startRow, GameBoard.startColumn}
		player.Direction = int(Up)
		fmt.Printf("--------- %s Finish 1 run, Rest %d Pieces. ---------\n", player.Name, player.Pieces)
	}

	if player.Pieces == 0 {
		fmt.Printf("||| --- %s Win --- |||\n", player.Name)
	} else {
		Set_Players(players)
	}
}

func main() {
	initGameBoard("GameBoard.in")

	for i := 0; i < PlayerNum; i++ {
		player := Player{Name: fmt.Sprintf("Player %d", i+1), Position: pos{1, 1}, Pieces: 4, Direction: int(Up), Turno: i}
		players = append(players, &player)
	}

	br := bufio.NewReader(os.Stdin)

	fmt.Print("Player n: ")
	srt, _ := br.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(srt))
	playern = num

	fmt.Print("Enter port: ")
	port, _ := br.ReadString('\n')
	port = strings.TrimSpace(port)
	Playerhost = fmt.Sprintf("localhost:%s", port)

	fmt.Print("Enter destiny: ")
	dest, _ := br.ReadString('\n')
	dest = strings.TrimSpace(dest)
	Destinyhost = fmt.Sprintf("localhost:%s", dest)

	ln, _ := net.Listen("tcp", Playerhost)

	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
