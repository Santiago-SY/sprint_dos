package proyecto

import (
	"sync"
)

type player struct {
	nombre string
	ELO    int
}

type nodo struct {
	Player *player // Puntero al dato real
	Next   *nodo   // Puntero al siguiente nodo
}

type Match struct {
	ID      int
	players []player
}

type Queue struct {
	head *nodo      // Primer lugar (El que sale)
	tail *nodo      // Ultimo lugar (El que entra)
	size int        // Para llevar la cuenta rapido
	mu   sync.Mutex // Seguridad
}

type lobby struct {
	mu      sync.Mutex
	matches map[int]Match
}
