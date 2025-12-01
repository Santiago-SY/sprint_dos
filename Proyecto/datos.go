package proyecto

import (
	"sync"
)

type player struct {
	nombre string
	ELO    int
}

type Match struct {
	ID      int
	players []player
}

type Queue struct {
	mu      sync.Mutex
	players []player
}

type lobby struct {
	mu      sync.Mutex
	matches map[int]Match
}
