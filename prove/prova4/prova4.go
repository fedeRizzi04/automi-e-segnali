package main

import (
	"fmt"
	"math"
)

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

// Funzione per trovare il punto nel piano corrispondente a una cella della matrice
func matrixToPoint(i, j, startX, startY, dirX, dirY int) (int, int) {
	return startX + j*dirX, startY + i*dirY
}

// Funzione per riempire la matrice booleana e determinare l'esistenza del percorso
func fillDP(startX, startY, endX, endY int, obstacles []Rectangle, n, m int, dirX, dirY int) bool {
	// Inizializza la matrice booleana
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, m)
	}

	// Determina l'ordine delle righe e colonne in base alla direzione
	rowStart, rowEnd, rowStep := 0, n, 1
	if dirY < 0 {
		rowStart, rowEnd, rowStep = n-1, -1, -1
	}

	colStart, colEnd, colStep := 0, m, 1
	if dirX < 0 {
		colStart, colEnd, colStep = m-1, -1, -1
	}

	// Calcola l'indice iniziale del punto di partenza
	startRow := (startY - startY) / dirY
	startCol := (startX - startX) / dirX
	dp[startRow][startCol] = true

	// Riempi la matrice
	for i := rowStart; i != rowEnd; i += rowStep {
		for j := colStart; j != colEnd; j += colStep {
			if isInObstacle(startX+j*dirX, startY+i*dirY, obstacles) {
				continue
			}

			fromAbove := i-rowStep >= 0 && i-rowStep < n && dp[i-rowStep][j]
			fromLeft := j-colStep >= 0 && j-colStep < m && dp[i][j-colStep]

			dp[i][j] = (fromAbove || fromLeft)
		}
	}

	// Calcola l'indice del punto di arrivo
	endRow := (endY - startY) / dirY
	endCol := (endX - startX) / dirX

	// Restituisci true se esiste un percorso libero
	return dp[endRow][endCol]
}

func main() {
	startX, startY := 0, 0
	endX, endY := 8, 8
	obstacles := []Rectangle{
		{x0: 3, y0: 3, x1: 4, y1: 4},
		//{x0: 8, y0: 8, x1: 9, y1: 9},
	}

	// Calcola le dimensioni della griglia
	n := int(math.Abs(float64(endY-startY))) + 1
	m := int(math.Abs(float64(endX-startX))) + 1

	// Determina la direzione del percorso
	dirX := 1
	if endX < startX {
		dirX = -1
	}
	dirY := 1
	if endY < startY {
		dirY = -1
	}

	// Verifica l'esistenza del percorso
	if fillDP(startX, startY, endX, endY, obstacles, n, m, dirX, dirY) {
		fmt.Println("SI")
	} else {
		fmt.Println("NO")
	}
}
