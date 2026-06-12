package concurrence

// SendValue envoie une valeur dans le channel
func SendValue(ch chan int, value int) {
	// TODO: Envoyer value dans ch
}

// SumParallel calcule la somme dans une goroutine
func SumParallel(a, b int) int {
	// TODO: Créer le channel
	// TODO: Lancer la goroutine
	// TODO: Lire et retourner le résultat
	return 0
}

// PingPong simule un échange entre deux channels
func PingPong(count int) ([]string, []string) {
	pings := make(chan string, count)
	pongs := make(chan string, count)
	
	var pingResults []string
	var pongResults []string

	// TODO: Goroutine 1 - Envoie "ping" count fois dans pings, puis ferme pings
	
	// TODO: Goroutine 2 - Reçoit de pings, ajoute à pingResults, envoie "pong" dans pongs
	// Puis ferme pongs quand fini.

	// TODO: Collecter les résultats de pongs dans pongResults
	
	return pingResults, pongResults
}
