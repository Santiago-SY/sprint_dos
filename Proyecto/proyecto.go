package main

import (
	"fmt"
	"math/rand/v2" // Asegúrate de estar usando Go 1.22+ para v2
	"time"
)

func main() {
	fmt.Println("🔥 --- INICIANDO SERVIDOR DE MATCHMAKING (Modo Infinito) --- 🔥")

	// 1. Inicializar Queue y Lobby
	// Es vital inicializar el mapa del Lobby con make() para evitar errores nulos
	q := &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
	l := &Lobby{
		Matches: make(map[int]Match),
	}

	// 2. Lanzar el Matchmaker en paralelo
	// Se ejecuta en su propio hilo (goroutine)
	go Matchmaker(q, l)
	fmt.Println("✅ Matchmaker corriendo en segundo plano...")

	// 3. Generador de Tráfico Infinito
	// Usamos 'i' para darles un número único en el nombre
	i := 1

	for {
		// A. Generar ELO Random (0 - 1000)
		eloRandom := rand.IntN(1000)

		// B. Crear el Jugador
		nuevoJugador := &Player{
			Nombre: fmt.Sprintf("Jugador_%d", i),
			ELO:    eloRandom,
			// TimeJoined se pone solo dentro de Enqueue
		}

		// C. Meter a la Cola
		q.Enqueue(nuevoJugador)

		// (Opcional) Log para ver el flujo. Si va muy rápido, puedes comentarlo.
		// fmt.Printf("-> %s (ELO: %d)\n", nuevoJugador.Nombre, nuevoJugador.ELO)

		// D. Simulación de Tráfico Realista
		// Hacemos que entre gente cada 20ms a 50ms.
		// Esto genera aprox. 20-30 jugadores por segundo.
		delay := time.Duration(rand.IntN(30)+20) * time.Millisecond
		time.Sleep(delay)

		i++ // Preparamos el contador para el siguiente
	}
}
