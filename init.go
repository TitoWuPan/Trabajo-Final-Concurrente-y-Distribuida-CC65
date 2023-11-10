package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type Direction int

var remotehost string
var playern int = 0

const (
	PlayerNum           = 4
	Up        Direction = 0
	Down      Direction = 1
	Left      Direction = 2
	Right     Direction = 3
)

type pos struct {
	i, j int
}

type Player struct {
	Name      string
	Position  pos
	Pieces    int
	Direction int
	Turno     int
}

var players []Player
var Destinyhost string

func Set_Players(players []Player) {
	con, _ := net.Dial("tcp", Destinyhost)

	arrBytesJson, _ := json.Marshal(players)
	strMsgJson := string(arrBytesJson)

	fmt.Fprintln(con, strMsgJson)
}

func main() {
	for i := 0; i < PlayerNum; i++ {
		player := Player{Name: fmt.Sprintf("Player %d", i+1), Position: pos{1, 1}, Pieces: 4, Direction: int(Up), Turno: i}
		players = append(players, player)
	}

	br := bufio.NewReader(os.Stdin)

	fmt.Print("Enter destiny: ")
	dest, _ := br.ReadString('\n')
	dest = strings.TrimSpace(dest)
	Destinyhost = fmt.Sprintf("localhost:%s", dest)

	Set_Players(players)
}
