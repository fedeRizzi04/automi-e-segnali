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

// Funzione per determinare le dimensioni della griglia
func gridDimensions(x1, y1, x2, y2 int) (int, int, int, int) {
	minX, maxX := min(x1, x2), max(x1, x2)
	minY, maxY := min(y1, y2), max(y1, y2)

	return minX, maxX, minY, maxY
}

// Funzione generica per riempire la matrice DP considerando entrambe le direzioni
func fillDP(dp [][]int, obstacles []Rectangle, startX, endX, stepX, startY, endY, stepY int) {
	n := len(dp)
	m := len(dp[0])

	for x := startX; x != endX; x += stepX {
		for y := startY; y != endY; y += stepY {
			conta++
			if isInObstacle(x, y, obstacles) {
				dp[x][y] = inf
				continue
			}

			if x-stepX >= 0 && x-stepX < n {
				dp[x][y] = min(dp[x][y], dp[x-stepX][y]+1)
			}
			if y-stepY >= 0 && y-stepY < m {
				dp[x][y] = min(dp[x][y], dp[x][y-stepY]+1)
			}
		}
	}
}

// Funzione principale
func existsPath(startX, startY, endX, endY int, obstacles []Rectangle) bool {
	// Determina le dimensioni della griglia
	minX, maxX, minY, maxY := gridDimensions(startX, startY, endX, endY)
	n := maxX - minX + 1
	m := maxY - minY + 1
	// Matrice DP inizializzata
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			conta++
			dp[i][j] = inf
		}
	}
	fmt.Println("Conta1 ", conta)
	conta = 0
	// Punto di partenza
	dp[startX-minX][startY-minY] = 0

	// Riempie la matrice in base alle direzioni ottimali
	fillDP(dp, obstacles, 0, n, 1, 0, m, 1)

	fmt.Println("Conta2 ", conta)

	// Verifica il punto di arrivo
	return dp[endX-minX][endY-minY] != inf
}

// Funzioni di supporto
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var conta int = 0

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

	// Verifica se esiste un percorso
	if existsPath(nx, ny, x, y, obstacles) {
		fmt.Println("SI")
	} else {
		fmt.Println("NO")
	}

}
