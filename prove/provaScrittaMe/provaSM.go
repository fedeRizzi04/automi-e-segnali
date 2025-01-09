package main

import "fmt"

type punto struct {
	x, y int
}

func aggiungiPerimetro(x0 int, y0 int, x1 int, y1 int) {
	//aggiunta dei lati orizzontali
	for i := x0; i <= x1; i++ {
		puntiConosciuti[punto{i, y0}] = "O"
		puntiConosciuti[punto{i, y1}] = "O"
	}

	//aggiunta dei lati verticali
	for i := y0; i <= y1; i++ {
		puntiConosciuti[punto{x0, i}] = "O"
		puntiConosciuti[punto{x1, i}] = "O"
	}

	// Costo: O(n+m) con n = x1-x0 e m = y1-y0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dimensioneTabella(punto1 punto, punto2 punto) (int, int) {
	minx, maxx := min(punto1.x, punto2.x), max(punto1.x, punto2.x)
	miny, maxy := min(punto1.y, punto2.y), max(punto1.y, punto2.y)

	return maxy - miny + 1, maxx - minx + 1
}

func signum(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func direzioniPossibili(partenza punto, arrivo punto) (int, int) {
	dx := arrivo.x - partenza.x
	dy := arrivo.y - partenza.y
	return signum(dx), signum(dy)

}

func inizioTabella(dx int, dy int, n int, m int) (int, int) {
	var startI, startJ int

	if dx == -1 {
		startJ = m - 1
	} else {
		startJ = 0
	}

	if dy == -1 {
		startI = n - 1
	} else {
		startI = 0
	}

	return startI, startJ
}

func normalizzaEntrata(i int, j int, partenza punto, arrivo punto, dx int, dy int, n int, m int) punto {
	var coordx, coordy int
	if dx == -1 {
		coordy = m - 1 - j + arrivo.y
	} else {
		coordy = j + partenza.y
	}

	if dy == -1 {
		coordx = n - 1 - i + arrivo.x
	} else {
		coordx = i + partenza.x
	}

	return punto{coordx, coordy}

}

/*
func esistePercorso(partenza punto, arrivo punto) bool {

	n, m := dimensioneTabella(partenza, arrivo)
	tabella := make([][]bool, n)
	for i := 0; i < n; i++ {
		tabella[i] = make([]bool, m)
	}

	dx, dy := direzioniPossibili(partenza, arrivo)
	// ora bisogna normalizzare i punti negli indici della tabella
	// questo varia in base a come si scorre la tabella (da sinistra a destra o viceversa) (o da sopra a sotto o viceversa)
	//se dx = 1 allora si scorre da sinistra a destra, altrimenti da destra a sinistra
	//se dy = 1 allora si scorre da sotto (basso indice y) a sopra (alto indice y), altrimenti da sopra a sotto
	// se dx = 1 e dy = 1 allora si scorre da sinistra a destra e da sotto a sopra
	// se dx = 1 e dy = -1 allora si scorre da sinistra a destra e da sopra a sotto
	// se dx = -1 e dy = 1 allora si scorre da destra a sinistra e da sotto a sopra
	// se dx = -1 e dy = -1 allora si scorre da destra a sinistra e da sopra a sotto

	//da fare

}
*/

var puntiConosciuti map[punto]string = make(map[punto]string)

func main() {

	partenza := punto{3, 8}
	arrivo := punto{8, 5}
	//CONTROLLARE DIMENSIONE TABELLA CON NUMERI NEGATIVI
	fmt.Println(dimensioneTabella(partenza, arrivo))
	fmt.Println(inizioTabella(1, -1, 4, 6))
	aggiungiPerimetro(2, 2, 6, 6)
	aggiungiPerimetro(5, 2, 6, 8)
	aggiungiPerimetro(4, 10, 12, 11)

	fmt.Println(normalizzaEntrata(3, 0, partenza, arrivo, 1, -1, 6, 4))
}
