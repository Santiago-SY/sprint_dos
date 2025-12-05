package main

import "time"

// encolar
//
//	Receptor      Nombre      Parametros
func (q *Queue) Enqueue(p *Player) {
	q.mu.Lock()
	defer q.mu.Unlock()
	p.EnCola = true
	p.TimeJoined = time.Now()
	nuevoNodo := &Nodo{
		Player: p,
		Next:   nil,
	}
	if q.head == nil {
		q.head = nuevoNodo
		q.tail = nuevoNodo
		q.size++
	} else {
		q.tail.Next = nuevoNodo
		q.tail = nuevoNodo
		q.size++
	}

}

// sacar de la cola
func (q *Queue) Dequeue() *Player {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil {
		return nil
	}
	jugadorsaliente := q.head.Player
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return jugadorsaliente
}

// sacar de la cola a alguien especifico
func (q *Queue) Remover_player(p *Player) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil {
		return
	}
	if q.head.Player == p {
		q.head = q.head.Next
		if q.head == nil {
			q.tail = nil
		}
		q.size--
		p.EnCola = false
		return
	}
	anterior := q.head
	actual := q.head.Next
	for actual != nil {
		if actual.Player == p {
			if actual == q.tail {
				q.tail = anterior
			}
			anterior.Next = actual.Next
			q.size--
			p.EnCola = false
			return
		}
		anterior = actual
		actual = actual.Next
	}
}

// ver el primero en la lista
func (q *Queue) Top() *Player {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head != nil {
		return q.head.Player
	}
	return nil
}

// EsVacia devuelve true si no hay nadie en la cola
func (q *Queue) EsVacia() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.head == nil
}

// Rotar mueve el primer jugador al final de la cola
func (q *Queue) Rotar() {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Si la cola está vacía o solo tiene 1 persona, no hacemos nada
	if q.head == nil || q.head.Next == nil {
		return
	}

	// Guardamos el nodo que vamos a mover (el actual primero)
	nodoAmover := q.head

	// El segundo pasa a ser el nuevo primero
	q.head = q.head.Next

	// El nodo movido se pega después del actual último
	q.tail.Next = nodoAmover

	// Actualizamos el puntero tail para que sea el nodo movido
	q.tail = nodoAmover

	// Importante: El nuevo último no debe apuntar a nadie
	q.tail.Next = nil
}

// queue.go

// Devuelve un slice con los primeros N jugadores (o menos si no hay tantos)
// Sin sacarlos de la cola (Peek multiple)
func (q *Queue) ObtenerPoolDeCandidatos(n int) []*Player {
	q.mu.Lock()
	defer q.mu.Unlock()

	candidatos := make([]*Player, 0, n)
	actual := q.head

	for actual != nil && len(candidatos) < n {
		candidatos = append(candidatos, actual.Player)
		actual = actual.Next
	}
	return candidatos
}
