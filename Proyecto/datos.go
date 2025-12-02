package proyecto

import (
	"sync"
	"time"
)

type Player struct {
	Nombre     string
	ELO        int
	IsPlaying  bool      // ¿Está ocupado? Evita que entre a la cola dos veces
	TimeJoined time.Time // Hora exacta a la que entró a la Queue (Para prioridad)
}

type Nodo struct {
	Player *Player // Puntero al dato real
	Next   *Nodo   // Puntero al siguiente nodo
}

type Match struct {
	ID      int
	Players []Player
}

type Queue struct {
	head *Nodo      // Primer lugar (El que sale)
	tail *Nodo      // Ultimo lugar (El que entra)
	size int        // Para llevar la cuenta rapido
	mu   sync.Mutex // Seguridad
}

type Lobby struct {
	mu      sync.Mutex
	Matches map[int]Match
}
