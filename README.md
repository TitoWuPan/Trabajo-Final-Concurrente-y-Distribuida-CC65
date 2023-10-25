# Trabajo Final 
Programación Concurrente y Distribuida

## Integrantes :
- Guillén Rojas, Daniel Carlos	-	U201920113
- Wu Pan, Tito Peng 	-		U201921200
- Sebastian Gonzales	-		U201923816

# Descripción del Juego:
El Ludo modificado es una versión ampliada y adaptada del popular juego de mesa Ludo. En esta versión, los jugadores compiten para llevar a sus personajes a través de un peligroso laberinto lleno de obstáculos y desafíos. Cada jugador tiene la tarea de guiar a sus personajes a través del laberinto y llegar a la meta antes que los demás. 

## Reglas del Juego:
### Tablero del Laberinto:
- El tablero está dividido en casillas con caminos, obstáculos.
- Cada jugador tiene cuatro personajes que comienzan en puntos de partida específicos.

### Turnos y Movimientos:
- Los jugadores se turnan para lanzar un dado y mover a sus personajes.
- Los jugadores lanzan tres dados, dos dados normales (del 1 al 6) y uno con la operación (sumar o restar) para determinar cuántos pasos pueden avanzar o retroceder en su turno.
- Los jugadores pueden mover un solo personaje por turno.
- Los personajes deben avanzar exactamente la cantidad de pasos indicados por la operación de los dados (valor del primer dado y operador (+ -) con el valor del segundo dado).

### Obstáculos:
- El laberinto está lleno de obstáculos como paredes, trampas y criaturas que bloquean el paso de los personajes en varias casillas.
- Si al personaje le toca avanzar hacia una casilla con obstáculo entonces el jugador pierde el turno y continua el siguiente jugador.
  
### Objetivo:
- El objetivo es llevar a los cuatro personajes desde los puntos de partida hasta la meta en el menor número de turnos posible.
- El primer jugador en llevar a todos sus personajes a la meta gana el juego.
  
### Modificaciones y Uso de Canales:
En esta versión modificada del juego, los jugadores y el tablero de juego están representados como entidades concurrentes separadas que se comunican a través de canales. Cada jugador tiene su propio canal de comunicación con el tablero del juego para enviar movimientos y recibir actualizaciones del estado del juego. 

