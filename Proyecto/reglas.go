package main

import (
	"fmt"
	"sort"
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
	// 1. ORDENAR: De Mayor a Menor ELO
	// La función anónima recibe i, j (índices) y nosotros comparamos los valores
	sort.Slice(jugadores, func(i, j int) bool {
		return jugadores[i].ELO > jugadores[j].ELO
	})

	// 2. INICIALIZAR EQUIPOS
	// Usamos make con []*Player (punteros)
	teamA := make([]*Player, 0)
	teamB := make([]*Player, 0)

	eloA := 0
	eloB := 0

	// 3. REPARTIR (Algoritmo Greedy)
	// Recorremos la lista ya ordenada jugador por jugador
	for _, p := range jugadores {

		// ¿Quién va perdiendo en suma de ELO? Le damos el siguiente jugador
		if eloA <= eloB {
			teamA = append(teamA, p)
			eloA += p.ELO
		} else {
			teamB = append(teamB, p)
			eloB += p.ELO
		}
	}

	return teamA, teamB
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
