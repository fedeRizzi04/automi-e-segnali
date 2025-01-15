package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type punto struct {
	x int
	y int
}

type ostacolo struct {
	bassoSx punto
	altoDx  punto
}

type ListOstacoli interface {
	AddOstacolo(r ostacolo)
	ContienePunto(p punto) bool
	StampaOstacoli()
}

type piano *struct {
	automi          map[string]punto
	ostacoli        ListOstacoli
	puntiConosciuti map[punto]string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var p piano
	for scanner.Scan() {
		str := scanner.Text()
		com := strings.Fields(str)[0]
		if com == "c" {
			p = newPiano()
		} else {
			esegui(p, str)
		}
	}
}

// definizione del tipo implementante l'interfaccia ListRettangoli
type NodoOstacolo struct {
	ret  *ostacolo
	next *NodoOstacolo
}
type myListOstacoli struct {
	head *NodoOstacolo
	tail *NodoOstacolo
}

func (l *myListOstacoli) AddOstacolo(r ostacolo) {
	n := &NodoOstacolo{&r, nil}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

func (l *myListOstacoli) ContienePunto(p punto) bool {
	for n := l.head; n != nil; n = n.next {
		if p.x >= n.ret.bassoSx.x && p.x <= n.ret.altoDx.x && p.y >= n.ret.bassoSx.y && p.y <= n.ret.altoDx.y {
			return true
		}
	}
	return false
}

func (l *myListOstacoli) StampaOstacoli() {
	for n := l.head; n != nil; n = n.next {
		fmt.Printf("(%d,%d)(%d,%d)\n", n.ret.bassoSx.x, n.ret.bassoSx.y, n.ret.altoDx.x, n.ret.altoDx.y)
	}
}

func esegui(p piano, s string) {
	comandi := strings.Fields(s)
	var a, b int
	if len(comandi) >= 3 {
		a, _ = strconv.Atoi(comandi[1])
		b, _ = strconv.Atoi(comandi[2])
	}

	switch comandi[0] {
	case "s":
		stato(p, a, b)
	case "S":
		stampa(p)
	case "a":
		automa(p, a, b, comandi[3])
	case "o":
		var c, d int
		c, _ = strconv.Atoi(comandi[3])
		d, _ = strconv.Atoi(comandi[4])
		aggiungiOstacolo(p, a, b, c, d)
	case "r":
		richiamo(p, a, b, comandi[3])
	case "p":
		posizioni(p, comandi[1])
	case "e":
		nomeAutoma := comandi[3]
		pAutoma, esiste := p.automi[nomeAutoma]
		if !esiste {
			fmt.Println("NO")
			break
		}
		p1, p2 := scegliPartenza(pAutoma, punto{a, b})
		if nomeAutoma != "" && esistePercorso(p, p1, p2) {
			fmt.Println("SI")
		} else {
			fmt.Println("NO")
		}
	case "f":
		os.Exit(0)
	default:
		fmt.Println("Comando non riconosciuto")
	}

}

func contiene(rett ostacolo, punto punto) bool {
	return (punto.x >= rett.bassoSx.x && punto.x <= rett.altoDx.x && punto.y >= rett.bassoSx.y && punto.y <= rett.altoDx.y)
}

func newPiano() piano {
	return &struct {
		automi          map[string]punto
		ostacoli        ListOstacoli
		puntiConosciuti map[punto]string
	}{
		automi:          make(map[string]punto),
		ostacoli:        &myListOstacoli{},
		puntiConosciuti: make(map[punto]string),
	}
}

func stampa(p piano) {
	fmt.Println("(")
	for k, v := range p.automi {
		fmt.Printf("%s: %d,%d\n", k, v.x, v.y)
	}
	fmt.Println(")")
	fmt.Println("[")
	p.ostacoli.StampaOstacoli()
	fmt.Println("]")
}

func presenzaErimuovi(p piano, n string) {
	puntoPossibile, ok := p.automi[n]
	if ok {
		quanti, _ := strconv.Atoi(p.puntiConosciuti[puntoPossibile][1:])
		quanti--
		if quanti == 0 {
			p.puntiConosciuti[puntoPossibile] = "E"
		} else {
			p.puntiConosciuti[puntoPossibile] = "A" + strconv.Itoa(quanti)
		}
	}
}

func automa(p piano, x int, y int, n string) {

	punto := punto{x, y}
	// controllo se si conosce già cosa c'è nel punto di arrivo
	str, esiste := p.puntiConosciuti[punto]
	if esiste {
		switch string(str[0]) {
		case "A":
			// se c'è un automa allora posso mettere sicuramente un automa nel punto di arrivo, aumentando il numero di automi in quel punto
			quanti, _ := strconv.Atoi(str[1:])
			quanti++
			p.puntiConosciuti[punto] = "A" + strconv.Itoa(quanti)
			// ora devo controllare se c'è un automa con lo stesso nome in un altro punto, se si devo rimuoverlo
			presenzaErimuovi(p, n)
			p.automi[n] = punto
		case "E":
			// se il punto di arrivo è vuoto allora posso mettere un automa in quel punto
			presenzaErimuovi(p, n)
			p.automi[n] = punto
			p.puntiConosciuti[punto] = "A1"
		}
		return // se c'è un ostacolo non faccio nulla
	}

	// se non si conosce cosa c'è nel punto di arrivo, allora bisogna controllare se c'è un ostacolo
	if !p.ostacoli.ContienePunto(punto) {
		presenzaErimuovi(p, n)
		p.automi[n] = punto
		p.puntiConosciuti[punto] = "A1"
	}
}

// dato un nuovo ostacolo aggiunto, aggiunge i punti del perimetro di tale ostacolo alla mappa dei punti conosciuti
func aggiungiPerimetro(p piano, x0 int, y0 int, x1 int, y1 int) {
	for i := x0; i <= x1; i++ {
		p.puntiConosciuti[punto{i, y0}] = "O"
		p.puntiConosciuti[punto{i, y1}] = "O"
	}

	for i := y0; i <= y1; i++ {
		p.puntiConosciuti[punto{x0, i}] = "O"
		p.puntiConosciuti[punto{x1, i}] = "O"
	}
}

func aggiungiOstacolo(p piano, x0 int, y0 int, x1 int, y1 int) {
	p1 := punto{x0, y0}
	p2 := punto{x1, y1}
	rett := ostacolo{p1, p2}
	for _, punto := range p.automi {
		if contiene(rett, punto) {
			return
		}
	}

	p.ostacoli.AddOstacolo(rett)
	aggiungiPerimetro(p, x0, y0, x1, y1)
}

func rangeAutomiPrefisso(p piano, s string, azione func(string, punto)) {
	for k, v := range p.automi {
		if strings.HasPrefix(k, s) {
			azione(k, v)
		}
	}
}

func posizioni(p piano, s string) {
	fmt.Println("(")
	rangeAutomiPrefisso(p, s, func(k string, v punto) {
		fmt.Printf("%s: %d,%d\n", k, v.x, v.y)
	})
	fmt.Println(")")
}

func automiPrefisso(p piano, s string) []string {
	nomiAutomi := []string{}
	rangeAutomiPrefisso(p, s, func(k string, p punto) {
		nomiAutomi = append(nomiAutomi, k)
	})
	return nomiAutomi
}

func distanzaPunti(p1 punto, p2 punto) int {
	return int(math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y)))
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

// dato un punto di arrivo e un automa in posizione, restituisce le direzioni in cui l'automa può muoversi in un percorso con lunghezza
// minima. La prima variabile di ritorno indica la direzione sull'asse x, la seconda sull'asse y (e tali valori possono essere -1, 0, 1)
func direzioni(puntoAutoma punto, puntoArrivo punto) (int, int) {
	dx := signum(puntoArrivo.x - puntoAutoma.x)
	dy := signum(puntoArrivo.y - puntoAutoma.y)
	return dx, dy
}

func dimensioneTabella(punto1 punto, punto2 punto) (int, int) {
	colonne := int(math.Abs(float64(punto2.x-punto1.x))) + 1
	righe := int(math.Abs(float64(punto2.y-punto1.y))) + 1
	return righe, colonne
}

// con l'assunzione che il punto di partenza debba avere direzione dx = 1 o 0, percorriamo la tabella sempre da sinistra verso destra
func esistePercorso(p piano, partenza punto, arrivo punto) bool {

	if partenza == arrivo { // il percorso non sarebbe di distanza al minimo 1
		return false
	}
	righe, colonne := dimensioneTabella(partenza, arrivo)
	_, diry := direzioni(partenza, arrivo)
	partenzaRighe, finaleRighe := partenzaFineRighe(diry, righe)

	dp := make([][]bool, righe)
	for i := range dp {
		dp[i] = make([]bool, colonne)
	}
	dp[partenzaRighe][0] = true
	if diry == 0 {
		diry = 1 //in questo modo non va in loop il ciclo for successivo
	}

	for i := partenzaRighe; i != finaleRighe; i += diry {
		for j := 0; j < colonne; j++ {
			// se il punto corrente è un ostacolo, non è raggiungibile
			if str := p.puntiConosciuti[normalizzaEntrataTabella(partenza, arrivo, diry, i, j)]; str == "O" {
				dp[i][j] = false
				continue
			}
			if partenzaRighe == i && j == 0 {
				continue
			}
			icontrollo := i - diry
			//controllo se il punto sulla riga precedente è raggiungibile da un percoro libero minimo a partire dal punto di arrivo. Nelle prossime righe non è possibile avere un outOfBound nella slice per via della lazy evaluation sulla condizione
			if icontrollo >= 0 && icontrollo < righe && dp[icontrollo][j] {
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

// data la direzione sull'asse y del percorso minimo, restituisce la riga di partenza della matrice dp da cui partire e quella da mettere nella condizione sul for
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

func richiamo(p piano, x int, y int, s string) {

	punto := punto{x, y}
	//controllo se si conosce già cosa se c'è un ostacolo in tale punto
	str := p.puntiConosciuti[punto]
	if str == "O" {
		return
	}
	//se non conosco cosa c'è devo vedere se c'è un ostacolo o meno
	if str == "" && p.ostacoli.ContienePunto(punto) {
		return
	}

	nomiAutomi := automiPrefisso(p, s)
	automiDaSpostare := []string{}
	min := -1
	for _, nome := range nomiAutomi {
		//devo mettere come primo punto quello che deve andare verso destra (con passi unitari orizzontali dove la x aumenta di 1)
		p1, p2 := scegliPartenza(p.automi[nome], punto)
		d := distanzaPunti(p1, p2)
		// se la distanza tra l'automa e il punto di richiamo è maggiore della distanza minima trovata finora, non ha senso calcolare l'esistenza di un percorso
		if min != -1 && d > min {
			continue
		}
		if esistePercorso(p, p1, p2) {
			if min == -1 || d < min {
				min = d
				automiDaSpostare = []string{nome}
			} else if d == min {
				automiDaSpostare = append(automiDaSpostare, nome)
			}
		}
	}
	// sposto gli automi che possono raggiungere il punto di richiamo
	spostaAutomi(automiDaSpostare, p, punto)
}

func scegliPartenza(p1 punto, p2 punto) (punto, punto) {
	dx, _ := direzioni(p1, p2)
	if dx == 1 {
		return p1, p2
	}
	return p2, p1
}

func spostaAutomi(automiDaSpostare []string, p piano, punto punto) {
	var quanti int = 0
	str := p.puntiConosciuti[punto]
	if str != "" && string(str[0]) == "A" {
		quanti, _ = strconv.Atoi(str[1:])
	}
	quanti += len(automiDaSpostare)

	for _, nome := range automiDaSpostare {
		presenzaErimuovi(p, nome)
		p.automi[nome] = punto
	}
	if quanti > 0 {
		p.puntiConosciuti[punto] = "A" + strconv.Itoa(quanti)
	}

}

func stato(p piano, x int, y int) {

	punto := punto{x, y}
	s := p.puntiConosciuti[punto]
	if s != "" {
		fmt.Println(string(s[0]))
		return
	}

	if p.ostacoli.ContienePunto(punto) {
		fmt.Println("O")
	} else {
		fmt.Println("E")
		p.puntiConosciuti[punto] = "E"
	}
}
