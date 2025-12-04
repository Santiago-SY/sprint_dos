package proyecto

import (
	"fmt"
	"math/rand/v2"
	"time"
)

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

func Matchmaker(q *Queue, l *Lobby) {
	for {
		if q.EsVacia() {
			time.Sleep(time.Second)
			continue
		}
		primero := q.Top()
		if primero == nil {
			continue
		}
		time.Since(q.Top().TimeJoined)
		rango := 50 + int(time.Second)*5
		equipo := encontrar_partida(q, primero, rango)
		if equipo != nil {
			nuevalobby :=  Match[nuevalobby, rand.IntN(1000)]
			for equipo != nil {
				nuevalobby.append(jugadorsaliente)
				q.Remover_player(jugadorsaliente)
			}
			l.AgregarMatch(nuevalobby)
			fmt.Println("Partida encontrada", nuevalobby.ID)
		}
		if equipo == nil {
			time.Sleep(time.Second)
		}
	}
}
