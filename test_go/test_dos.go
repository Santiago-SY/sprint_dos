package main

import (
	"fmt"
	"sync"
)

type Patio struct {
	mu       sync.Mutex
	ocupante string
}

func entrar(patio *Patio, nombre string, wg *sync.WaitGroup) {
	patio.mu.Lock()
	defer patio.mu.Unlock()
	defer wg.Done()
	patio.ocupante = nombre
	fmt.Println("el patio esta siendo usado por ", nombre)

	//falta return
}

func main() {
	p := &Patio{}
	var wg sync.WaitGroup
	wg.Add(2)
	go entrar(p, "Perro", &wg)
	go entrar(p, "Gato", &wg)
}
