# Trabajo Final 
Programación Concurrente y Distribuida

## Integrantes :
- Guillén Rojas, Daniel Carlos	-	U201920113
- Wu Pan, Tito Peng 	-		U201921200
- Sebastian Gonzales	-		U201923816

# Descripción del Juego:
El Ludo modificado es una versión ampliada y adaptada del popular juego de mesa Ludo. En esta versión, los jugadores compiten para llevar a sus personajes a través de un peligroso laberinto lleno de obstáculos y desafíos. Cada jugador tiene la tarea de guiar a sus personajes a través del laberinto y llegar a la meta antes que los demás. 

## Tareas a Implementar:
### Implementación de Canales:
- Cada jugador tiene un canal de comunicación bidireccional con el tablero.
- Los jugadores enviarán los resultados de los lanzamientos de dados y recibirán instrucciones para mover a sus personajes.
- 
### Lógica del Movimiento:
- Implementar la lógica para mover los personajes en el laberinto según las reglas del juego.
- Manejar las interacciones con obstáculos.

### Sincronización y Gestión de Turnos:
- Garantizar que los movimientos de los jugadores se manejen de manera sincronizada y que un jugador no pueda moverse fuera de su turno.
- Implementar la gestión de turnos para alternar entre los jugadores.

### Fin del Juego:
- Detectar cuando un jugador ha llevado a todos sus personajes a la meta y declarar a ese jugador como ganador.
- Recuerda considerar la concurrencia y la sincronización al implementar las interacciones a través de los canales para garantizar que el juego funcione de manera justa y sin errores durante las operaciones simultáneas de los jugadores. 


