package main

import (
	"fmt"
	"math"
)

const inf = math.MaxInt32

// Tipo per rappresentare un ostacolo come un rettangolo
type Rectangle struct {
	x0, y0, x1, y1 int
}

// Funzione per controllare se un punto è all'interno di un ostacolo
func isInObstacle(x, y int, obstacles []Rectangle) bool {
	for _, rect := range obstacles {
		if x >= rect.x0 && x <= rect.x1 && y >= rect.y0 && y <= rect.y1 {
			return true
		}
	}
	return false
}

func distanza(x0, y0, x1, y1 int) int {
	return int(math.Abs(float64(x1-x0)) + math.Abs(float64(y1-y0)))
}

// Funzione principale
func existsPath(nx, ny, x, y int, obstacles []Rectangle, n, m int) bool {
	// Normalizza le coordinate per lavorare con indici non negativi
	fmt.Println(distanza(nx, ny, x, y))
	minX, minY := min(nx, x), min(ny, y)
	nx, ny = nx-minX, ny-minY
	x, y = x-minX, y-minY

	// Dimensioni della griglia normalizzata
	n, m = max(nx, x)+1, max(ny, y)+1
	conta := 0
	// Matrice DP
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			conta++
			dp[i][j] = inf // Inizializza a infinito
		}
	}
	fmt.Println("Conta1 ", conta)
	conta = 0
	// Punto di partenza
	dp[nx][ny] = 0

	// Primo passaggio: verso giù e verso destra (scorrimento normale)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			conta++
			if isInObstacle(i+minX, j+minY, obstacles) {
				continue // Salta i punti negli ostacoli
			}

			// Sposta verso il basso
			if i+1 < n && !isInObstacle(i+1+minX, j+minY, obstacles) && dp[i][j] != inf {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
			}

			// Sposta verso destra
			if j+1 < m && !isInObstacle(i+minX, j+1+minY, obstacles) && dp[i][j] != inf {
				dp[i][j+1] = min(dp[i][j+1], dp[i][j]+1)
			}
		}
	}

	// Secondo passaggio: verso l'alto e verso sinistra (scorrimento inverso)
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			conta++
			if isInObstacle(i+minX, j+minY, obstacles) {
				continue // Salta i punti negli ostacoli
			}

			// Sposta verso il basso (già controllato nel primo passaggio)
			if i-1 >= 0 && !isInObstacle(i-1+minX, j+minY, obstacles) && dp[i][j] != inf {
				dp[i-1][j] = min(dp[i-1][j], dp[i][j]+1)
			}

			// Sposta verso destra (già controllato nel primo passaggio)
			if j-1 >= 0 && !isInObstacle(i+minX, j-1+minY, obstacles) && dp[i][j] != inf {
				dp[i][j-1] = min(dp[i][j-1], dp[i][j]+1)
			}
		}
	}

	fmt.Println("Conta2 ", conta)

	// Se il valore di dp[x][y] è infinito, il percorso non esiste
	return dp[x][y] != inf
}

// Funzione di supporto per il massimo tra due valori
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Funzione di supporto per il minimo tra due valori
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Esempio d'uso
func main() {
	// Coordinate di partenza e arrivo
	nx, ny := 3, 8
	x, y := 8, 5

	// Ostacoli
	obstacles := []Rectangle{
		{x0: 2, y0: 2, x1: 6, y1: 6},
		{x0: 5, y0: 2, x1: 6, y1: 8},
		{x0: 4, y0: 10, x1: 12, y1: 11},
	}

	// Dimensioni della griglia
	n, m := 6, 4

	// Verifica se esiste un percorso
	if existsPath(nx, ny, x, y, obstacles, n, m) {
		fmt.Println("SI")
	} else {
		fmt.Println("NO")
	}

	// Coordinate di partenza e arrivo
	nx, ny = 2, 1
	x, y = 5, 3

	// Ostacoli
	obstacles = []Rectangle{
		{x0: 1, y0: 3, x1: 4, y1: 6},
		{x0: 3, y0: 2, x1: 4, y1: 7},
		{x0: 3, y0: 0, x1: 15, y1: 1},
	}

	// Dimensioni della griglia
	n, m = 5, 2

	// Verifica se esiste un percorso
	if existsPath(nx, ny, x, y, obstacles, n, m) {
		fmt.Println("SI")
	} else {
		fmt.Println("NO")
	}
}
