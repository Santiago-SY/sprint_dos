package proyecto

func (l *Lobby) AgregarMatch(m Match) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Matches[m.ID] = m

}

func (l *Lobby) RemoverMatch(ID int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.Matches, ID)
}
