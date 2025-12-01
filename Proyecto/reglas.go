package proyecto

import (
	"fmt"
	"sync"
)

func Add_to_Queue(player player, queue *Queue, wg *sync.WaitGroup) {
	queue.mu.Lock()
	defer queue.mu.Unlock()
	defer wg.Done()
	queue.players = append(queue.players, player)
	fmt.Println("el jugador", player.nombre, player.ELO, "ELO esta en queue")
}
