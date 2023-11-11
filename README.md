# Tarea Académica 4 : 2023-2
Programación Concurrente y Distribuida

## Integrantes :
- Guillén Rojas, Daniel Carlos	-	U201920113
- Wu Pan, Tito Peng 	-		U201921200
- Sebastian Gonzales	-		U201923816

## Objetivos
- El objetivo del proyecto es simular el juego de niños Ludo modificado usando programación concurrente, canales y algoritmos distribuidos.
- La simulación debe correr concurrentemente usando algoritmo distribuido y manejar un grupo grande de jugadores como host donde la comunicación es a través de puertos y sincronización usando canales.
- La simulación debe mostrar el progreso del juego en tiempo real. 

## Desarrollo
- Para poder dar solución al primer objetivo se creó la función ```Set_Players()``` y la función ```handle()```

  La función ```Set_Players()``` toma un arreglo de punteros a Player, lo convierte a formato JSON y luego envía ese JSON a través de una conexión TCP al destino especificado por Destinyhost.
  ```ruby
  func Set_Players(players []*Player) {
	con, _ := net.Dial("tcp", Destinyhost)

	arrBytesJson, _ := json.Marshal(players)
	strMsgJson := string(arrBytesJson)

	fmt.Fprintln(con, strMsgJson)
  }
  ```
  La función ```handle()``` se encarga de recibir la información del jugador a través de una conexión TCP, realizar su movimiento en el juego, actualizar el estado y enviar la información actualizada a través de la conexión. Este ciclo continúa mientras el jugador tenga piezas restantes en el juego.
  ```ruby
  func handle(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	str, _ := br.ReadString('\n')

	json.Unmarshal([]byte(str), players)

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
  ```
- El desarrollo del segundo objetivo está implementado en la funciión ```main()```

  La función ```main()``` establece el entorno inicial del juego, crea jugadores, maneja la entrada del jugador actual, inicia un servidor TCP para escuchar conexiones entrantes y luego, en un bucle infinito, acepta conexiones de jugadores y delega el manejo de cada conexión a la función ```handle()```.
  ```ruby
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
  ```
  - Por último, para validar el tercer objetivo, se desarrolló la funcionalidad del juego de la siguiente manera:
 
    Primero debemos validar en diferentes terminales el número de jugador ```Player n```, el puerto de cada jugador ```Enter port``` y por último el destino al cual deben ir ```Enter destiny```
    
    ![image](https://github.com/TitoWuPan/Trabajo-Final-Concurrente-y-Distribuida-CC65/assets/103372071/929d4515-6d08-409d-ae03-07d810764a31)
    
    Inicializamos el juego poniendo el punto de destino ```Enter destiny```
    
    ![image](https://github.com/TitoWuPan/Trabajo-Final-Concurrente-y-Distribuida-CC65/assets/103372071/5744647d-2763-45ca-8e4a-20b2c2fd68ad)
    
    Una vez inicializado los terminales de los jugadores nos mostrarán como fue su recorrido y cual es el ganador (en este caso el Player 2)
    
    ![image](https://github.com/TitoWuPan/Trabajo-Final-Concurrente-y-Distribuida-CC65/assets/103372071/9131a4a4-bb91-4acc-abc3-d94dbfca1d80)


    


