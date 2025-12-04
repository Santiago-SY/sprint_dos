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
