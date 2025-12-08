package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"sync/atomic"
	"time"
)

var idCounter int64

const MaxPartidas = 10

// reglas.go

func encontrar_partida(q *Queue, referencia *Player, rango int) []*Player {
	q.mu.Lock()
	defer q.mu.Unlock()

	equipo := make([]*Player, 0, 10)
	equipo = append(equipo, referencia)

	// Empezamos a buscar desde el principio de la cola (head)
	// No desde head.Next, porque referencia podría estar en el medio.
	actual := q.head

	for actual != nil && len(equipo) < 10 {

		// 🚨 FIX: Evitar agregarse a uno mismo
		if actual.Player == referencia {
			actual = actual.Next
			continue
		}

		diferencia := actual.Player.ELO - referencia.ELO
		if diferencia < 0 {
			diferencia = -diferencia
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
	// 1. ORDENAR (Igual que antes)
	sort.Slice(jugadores, func(i, j int) bool {
		return jugadores[i].ELO > jugadores[j].ELO
	})

	teamA := make([]*Player, 0)
	teamB := make([]*Player, 0)

	eloA := 0
	eloB := 0

	// 2. REPARTIR CON LÍMITE DE CUPO
	for _, p := range jugadores {

		// CASO 1: Equipo A está lleno (5), va forzado al B
		if len(teamA) == 5 {
			teamB = append(teamB, p)
			eloB += p.ELO
			continue
		}

		// CASO 2: Equipo B está lleno (5), va forzado al A
		if len(teamB) == 5 {
			teamA = append(teamA, p)
			eloA += p.ELO
			continue
		}

		// CASO 3: Ambos tienen lugar, aplicamos lógica Greedy normal
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

func AplicarResultados(ganadores []*Player, perdedores []*Player) {
	// Sumar a los ganadores
	for _, p := range ganadores {
		eloAnterior := p.ELO
		p.ELO += 25
		fmt.Printf("   📈 [WIN] %s | %d -> %d (+25)\n", p.Nombre, eloAnterior, p.ELO)
	}

	// Restar a los perdedores
	for _, p := range perdedores {
		eloAnterior := p.ELO
		p.ELO -= 25
		if p.ELO < 0 {
			p.ELO = 0 // Evitamos ELO negativo
		}
		fmt.Printf("   📉 [LOSE] %s | %d -> %d (-25)\n", p.Nombre, eloAnterior, p.ELO)
	}
}

func SimularPartida(m Match, l *Lobby) {
	// 1. Simular tiempo de juego
	//time.Sleep(3 * time.Second)

	// --- OPCIÓN B: MODO PRESENTACIÓN ---
	time.Sleep(5 * time.Second)

	// 2. Decidir Ganador (50/50)
	// rand.IntN(2) devuelve 0 o 1
	ganaA := rand.IntN(2) == 0

	fmt.Printf("\n[GAME] 🏆 Partida %d Finalizada.\n", m.ID)

	if ganaA {
		fmt.Println("       Resultado: Gana EQUIPO A 🔵")
		AplicarResultados(m.TeamA, m.TeamB)
	} else {
		fmt.Println("       Resultado: Gana EQUIPO B 🔴")
		AplicarResultados(m.TeamB, m.TeamA)
	}

	// 3. LIBERAR JUGADORES (Vuelta al ruedo)
	// Unimos los dos equipos en una sola lista temporal para recorrerlos rápido
	todos := append(m.TeamA, m.TeamB...)

	for _, p := range todos {
		p.EnPartida = false // ¡Libres! El Main ya los puede volver a elegir.
	}

	// 4. Limpiar el Lobby (Ya terminó)
	l.RemoverMatch(m.ID)

	fmt.Println("------------------------------------------------------")
}

// reglas.go

func Matchmaker(q *Queue, l *Lobby) {
	for {
		// 1. CHEQUEO DE HARDWARE (Límite de 10 partidas)
		activas := l.CantidadActivas() // Guardamos el número en una variable
		if l.CantidadActivas() >= MaxPartidas {
			// DEBUG
			fmt.Printf("⚠️ Servidores Llenos (%d/%d). Cola: %d esperando...\n", activas, MaxPartidas, q.size)
			time.Sleep(1 * time.Second)
			continue
		}

		if q.EsVacia() {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// 2. ESTRATEGIA WINDOWING:
		// Traemos a los 5 primeros de la fila para probar suerte con ellos.
		candidatos := q.ObtenerPoolDeCandidatos(5)
		matchFound := false

		for _, candidato := range candidatos {
			// Si el candidato ya está en partida (por algún race condition raro), saltarlo
			if candidato.EnPartida {
				continue
			}

			// Calculamos su rango personal
			tiempoEspera := time.Since(candidato.TimeJoined)
			rango := 50 + (int(tiempoEspera.Seconds()) * 25)

			// Buscamos equipo PARA ESTE candidato
			equipo := encontrar_partida(q, candidato, rango)

			if equipo != nil {
				// ¡EUREKA! Encontramos match (quizás no era el 1º de la fila, pero sirve)
				matchFound = true

				// Marcar y sacar de la cola
				for _, jugador := range equipo {
					jugador.EnPartida = true
					q.Remover_player(jugador)
				}

				// Crear partida
				teamA, teamB := BalancearEquipos(equipo)
				nuevoID := atomic.AddInt64(&idCounter, 1)
				nuevaPartida := Match{ID: int(nuevoID), TeamA: teamA, TeamB: teamB}
				l.AgregarMatch(nuevaPartida)

				// Logs
				avgA := 0
				for _, p := range teamA {
					avgA += p.ELO
				}
				if len(teamA) > 0 {
					avgA /= len(teamA)
				}
				avgB := 0
				for _, p := range teamB {
					avgB += p.ELO
				}
				if len(teamB) > 0 {
					avgB /= len(teamB)
				}

				fmt.Printf("[MATCH] ⚔️ Partida %d (Rango: %d) | 🔵 Avg: %d vs 🔴 Avg: %d\n", nuevoID, rango, avgA, avgB)
				go SimularPartida(nuevaPartida, l)

				// Importante: Si encontramos partida, dejamos de buscar en este ciclo
				// y volvemos a empezar para refrescar la lista.
				break
			}
		}

		// 3. GESTIÓN DE FRUSTRACIÓN
		// Si revisamos a los 5 primeros y NINGUNO pudo armar partida...
		if !matchFound {
			// Entonces sí, el primero es un "tapón" muy difícil. Lo rotamos.
			q.Rotar()
			// Pequeña pausa para no quemar CPU si nadie encuentra nada
			time.Sleep(10 * time.Millisecond)
		}
	}
}
