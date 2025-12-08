package main

import (
	"fmt"
	"math/rand/v2" // Si usas Go < 1.22 usa "math/rand"
	"time"
)

func main() {
	fmt.Println("🔥 --- INICIANDO SERVIDOR DE MATCHMAKING (Modo Persistente) --- 🔥")

	// 1. Inicializar Queue y Lobby
	q := &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
	// Importante: Inicializar el mapa
	l := &Lobby{
		Matches: make(map[int]Match),
	}

	// 2. Lanzar el Matchmaker (Cerebro)
	go Matchmaker(q, l)
	fmt.Println("✅ Matchmaker corriendo en segundo plano...")

	// 3. GENERAR POBLACIÓN (BASE DE DATOS EN MEMORIA)
	// Creamos 100 jugadores fijos que vivirán en el servidor
	poblacionTotal := 1000
	jugadoresDB := make([]*Player, poblacionTotal)

	fmt.Printf("👥 Generando población de %d jugadores...\n", poblacionTotal)
	for i := 0; i < poblacionTotal; i++ {
		jugadoresDB[i] = &Player{
			Nombre:    fmt.Sprintf("Player_%d", i+1),
			ELO:       rand.IntN(1000), // ELO Inicial aleatorio
			EnCola:    false,
			EnPartida: false,
		}
	}
	fmt.Println("✅ Población cargada. Iniciando simulación de tráfico...")
	fmt.Println("---------------------------------------------------------")

	// 4. BUCLE DE VIDA (Simulador de Tráfico)
	// Ya no creamos gente nueva. Elegimos gente existente para que juegue.
	for {
		// A. Elegir un jugador al azar de la población
		indiceRandom := rand.IntN(poblacionTotal)
		p := jugadoresDB[indiceRandom]

		// B. Intentar meterlo a jugar
		// Solo entra si NO está en la cola Y NO está jugando
		if !p.EnCola && !p.EnPartida {

			// Log de Login (con etiqueta [LOGIN])
			fmt.Printf("[LOGIN] 👤 %s (ELO: %d) entra a la cola.\n", p.Nombre, p.ELO)

			q.Enqueue(p)
		}

		// C. Velocidad del servidor
		// Hacemos que intenten entrar muy rápido (cada 10-50ms) para saturar la cola
		// y ver cómo el matchmaker trabaja bajo presión.
		//Opcion A = Performance
		//time.Sleep(time.Duration(rand.IntN(10)+1) * time.Millisecond)
		//Opcion B= Presentacion
		time.Sleep(time.Duration(rand.IntN(500)+200) * time.Millisecond)
	}
}
