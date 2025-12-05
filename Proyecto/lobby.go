package main

func (l *Lobby) AgregarMatch(m Match) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Matches[m.ID] = m

}

// Devuelve cuántas partidas hay jugando actualmente
func (l *Lobby) CantidadActivas() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.Matches)
}

func (l *Lobby) RemoverMatch(ID int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.Matches, ID)
}
