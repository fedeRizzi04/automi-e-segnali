package main

import (
	"math"
)

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

	colonne := int(math.Abs(float64(maxx-minx)) + 1)
	righe := int(math.Abs(float64(maxy-miny)) + 1)
	return righe, colonne

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

func direzioni(partenza punto, arrivo punto) (int, int) {
	dx := arrivo.x - partenza.x
	dy := arrivo.y - partenza.y
	return signum(dx), signum(dy)
}

var puntiConosciuti map[punto]string = make(map[punto]string)

// con l'assunzione che il punto di partenza debba avere direzione dx = 1 o 0, percorriamo la tabella sempre da sinistra verso destra
func esistePercorso(partenza punto, arrivo punto) bool {

	righe, colonne := dimensioneTabella(partenza, arrivo)
	_, diry := direzioni(partenza, arrivo)
	partenzaRighe, finaleRighe := partenzaFineRighe(diry, righe) // in base alla direzione di y, determiniamo la riga di partenza e di arrivo. Se la direzione va verso il basso, si parte dalla riga più in alto, se no dalla riga più in basso

	dp := make([][]bool, righe)
	for i := range dp {
		dp[i] = make([]bool, colonne)
	}
	dp[partenzaRighe][0] = true
	if diry == 0 {
		diry = 1
	}

	for i := partenzaRighe; i != finaleRighe; i += diry {
		for j := 0; j < colonne; j++ {
			// se il punto corrente è un ostacolo, non è raggiungibile
			if _, ok := puntiConosciuti[normalizzaEntrataTabella(partenza, arrivo, diry, i, j)]; ok {
				dp[i][j] = false
				continue
			}
			//per il punto di partenza non è necessario controllare se è raggiungibile da un percorso libero minimo a partire da esso, è già vero
			if partenzaRighe == i && j == 0 {
				continue
			}
			icontrollo := i - diry
			//controllo se il punto sulla riga precedente è raggiungibile da un percoro libero minimo a partire dal punto di arrivo. Nelle prossime righe non è possibile avere un outOfBound nella slice per via della lazy evaluation sulla condizione
			if icontrollo >= 0 && icontrollo < righe && icontrollo != i && dp[icontrollo][j] {
				dp[i][j] = true
				//se no provo col punto sulla colonna precedente
			} else if j-1 >= 0 && j-1 < colonne && dp[i][j-1] {
				dp[i][j] = true
			} else {
				//se non è raggiungibile da nessuno dei due punti precedenti, allora non è raggiungibile
				dp[i][j] = false
			}

		}
	}
	finaleRighe = rigaRitornoDP(finaleRighe, righe)
	return dp[finaleRighe][colonne-1]
}

func partenzaFineRighe(diry int, righe int) (int, int) {
	var partenzaI int
	var finaleI int
	if diry == -1 {
		partenzaI = righe - 1
		finaleI = -1
	} else {
		partenzaI = 0
		finaleI = righe
	}
	return partenzaI, finaleI
}

func rigaRitornoDP(finaleI int, righe int) int {
	if finaleI == -1 {
		finaleI = 0
	} else {
		finaleI = righe - 1
	}
	return finaleI
}

func normalizzaEntrataTabella(partenza punto, arrivo punto, dy int, i int, j int) punto {
	xPunto := 0
	yPunto := 0
	if dy == 1 {
		xPunto = j + partenza.x
		yPunto = i + partenza.y
	} else {
		xPunto = j + partenza.x
		yPunto = i + arrivo.y
	}
	return punto{xPunto, yPunto}
}

func main() {

	partenza := punto{1, 2}
	arrivo := punto{10, 2}

	//aggiungiPerimetro(1, 3, 4, 6)
	//aggiungiPerimetro(3, 2, 4, 7)
	//aggiungiPerimetro(3, 0, 15, 1)
	//aggiungiPerimetro(6, 1, 10, 3)
	//aggiungiPerimetro(6, 5, 7, 8)

	if esistePercorso(partenza, arrivo) {
		println("SI")
	} else {
		println("NO")
	}
}
