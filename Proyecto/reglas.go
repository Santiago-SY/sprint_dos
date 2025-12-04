package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var idCounter int64

func encontrar_partida(q *Queue, referencia *Player, rango int) []*Player {
	q.mu.Lock()
	defer q.mu.Unlock()

	equipo := make([]*Player, 0, 10)
	equipo = append(equipo, referencia)
	actual := q.head.Next

	for actual != nil && len(equipo) < 10 {

		diferencia := actual.Player.ELO - referencia.ELO
		if diferencia < 0 {
			diferencia = -diferencia // Hacemos positivo el valor
		}

		if diferencia <= rango {

			equipo = append(equipo, actual.Player)
		}

		actual = actual.Next
	}
	if len(equipo) == 10 {
		return equipo
	}
	return nil

}

func BalancearEquipos(jugadores []*Player) ([]*Player, []*Player) {
	if jugadores != nil {
		//	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
		sort.Slice(jugadores, func(jugadores[i].ELO, jugadores[j].ELO []*Player) bool {jugadores[i].ELO > jugadores[j].ELO})
		teamA : make([]Player, 0, 5)
		TeamB : make([]Player, 0, 5)
		teamA.ELO = 0
		teamB.ELO = 0
	}
}

func Matchmaker(q *Queue, l *Lobby) {
	for {
		if q.EsVacia() {
			time.Sleep(1 * time.Second)
			continue
		}
		primero := q.Top()
		if primero == nil {
			time.Sleep(1 * time.Second)
			continue
		}
		tiempoEspera := time.Since(primero.TimeJoined)
		segundos := int(tiempoEspera.Seconds())
		rango := 50 + (segundos * 5)
		equipo := encontrar_partida(q, primero, rango)
		if equipo != nil {
			nuevoID := atomic.AddInt64(&idCounter, 1)
			nuevaPartida := Match{
				ID:      int(nuevoID),
				Players: make([]Player, 0, 10),
			}
			for _, jugador := range equipo {
				jugador.EnPartida = true
				q.Remover_player(jugador)
				nuevaPartida.Players = append(nuevaPartida.Players, *jugador)
			}
			l.AgregarMatch(nuevaPartida)
			fmt.Println("Partida encontrada", nuevaPartida.ID)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
