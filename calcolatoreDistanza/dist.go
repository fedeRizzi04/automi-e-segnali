package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type punto struct {
	x, y int
}

func distanza(p1 punto, p2 punto) int {
	return int(math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y)))
}

func main() {

	x1, _ := strconv.Atoi(os.Args[1])
	y1, _ := strconv.Atoi(os.Args[2])
	x2, _ := strconv.Atoi(os.Args[3])
	y2, _ := strconv.Atoi(os.Args[4])

	p1 := punto{x1, y1}
	p2 := punto{x2, y2}
	fmt.Println(distanza(p1, p2))

}
