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

type rettangolo struct {
	bassoSx punto // Punto in basso a sinistra
	altoDx  punto // Punto in alto a destra
}

// lista di rettangoli
type ListRettangoli interface {

	// Aggiunge un rettangolo alla lista
	AddRettangolo(r rettangolo)
	// Dato un punto restituisce true se il punto è contenuto in almeno un rettangolo della lista
	ContienePunto(p punto) bool
	//stampa i rettangoli
	StampaRettangoli()
}

type piano *struct {
	automi   map[string]punto
	ostacoli ListRettangoli
	// Mappa che associa ai punti del piano visitati almeno una volta il relativo contenuto: A se contiene un automa, O se contiene un ostacolo, E se è vuoto
	puntiConosciuti map[punto]string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	piano := newPiano()
	/*
		piano := crea()
		automa(piano, 2, 1, "1")
		automa(piano, 2, 8, "10")
		automa(piano, 5, 11, "11")
		automa(piano, 8, 9, "100")
		automa(piano, 10, 6, "101")
		automa(piano, 12, 2, "111")
		automa(piano, 13, 3, "110")

		ostacolo(piano, 1, 3, 4, 6)
		ostacolo(piano, 3, 2, 4, 7)
		ostacolo(piano, 3, 0, 15, 1)
		ostacolo(piano, 6, 5, 7, 8)
		ostacolo(piano, 6, 1, 10, 3)

		stampa(piano)

		//fmt.Println(esistePercorso(piano, 5, 3, "100"))
		richiamo(piano, 5, 3, "1")
		stampa(piano)
	*/

	for scanner.Scan() {
		esegui(piano, scanner.Text())
	}

}

// definizione del tipo implementante l'interfaccia ListRettangoli
type NodoRett struct {
	ret  *rettangolo
	next *NodoRett
}
type myListRettangoli struct {
	head *NodoRett
	tail *NodoRett
}

func (l *myListRettangoli) AddRettangolo(r rettangolo) {
	n := &NodoRett{&r, nil}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

func (l *myListRettangoli) ContienePunto(p punto) bool {
	for n := l.head; n != nil; n = n.next {
		//da ottimizzare
		if p.x >= n.ret.bassoSx.x && p.x <= n.ret.altoDx.x && p.y >= n.ret.bassoSx.y && p.y <= n.ret.altoDx.y {
			return true
		}
	}
	return false
}

func (l *myListRettangoli) StampaRettangoli() {
	for n := l.head; n != nil; n = n.next {
		fmt.Printf("(%d,%d)(%d,%d)\n", n.ret.bassoSx.x, n.ret.bassoSx.y, n.ret.altoDx.x, n.ret.altoDx.y)
	}
}

// Interfaccia coda per la visita in ampiezza
type Coda interface {
	isEmpty() bool
	enqueue(p punto, distanza int)
	dequeue() (punto, int)
}

type NodoCoda struct {
	punto    punto
	distanza int
	next     *NodoCoda
}

type MyCoda struct {
	head *NodoCoda
	tail *NodoCoda
}

func (c *MyCoda) isEmpty() bool {
	return c.head == nil
}

func (c *MyCoda) enqueue(p punto, distanza int) {
	n := &NodoCoda{p, distanza, nil}
	if c.head == nil {
		c.head = n
	} else {
		c.tail.next = n
	}
	c.tail = n
}

func (c *MyCoda) dequeue() (punto, int) {
	p := c.head.punto
	d := c.head.distanza
	c.head = c.head.next
	if c.head == nil {
		c.tail = nil
	}
	return p, d
}

// FUNZIONI PER SPECIFICHE DI PROGETTAZIONE

func esegui(p piano, s string) {
	comandi := strings.Fields(s)
	var a, b int
	if len(comandi) >= 3 {
		a, _ = strconv.Atoi(comandi[1])
		b, _ = strconv.Atoi(comandi[2])
	}

	switch comandi[0] {
	case "c":
		crea(p)
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
		ostacolo(p, a, b, c, d)
	case "r":
		richiamo(p, a, b, comandi[3])
	case "p":
		posizioni(p, comandi[1])
	case "e":
		nomeAutoma := comandi[3]
		pAutoma := p.automi[nomeAutoma]
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

func contiene(rett rettangolo, punto punto) bool {
	return (punto.x >= rett.bassoSx.x && punto.x <= rett.altoDx.x && punto.y >= rett.bassoSx.y && punto.y <= rett.altoDx.y)

	// Costo: O(1)
}

func newPiano() piano {
	return &struct {
		automi          map[string]punto
		ostacoli        ListRettangoli
		puntiConosciuti map[punto]string
	}{
		automi:          make(map[string]punto),
		ostacoli:        &myListRettangoli{},
		puntiConosciuti: make(map[punto]string),
	}
}

func crea(p piano) {
	p.automi = make(map[string]punto)
	p.ostacoli = &myListRettangoli{}
	p.puntiConosciuti = make(map[punto]string)
	// Costo: O(1)
}

func stampa(p piano) {

	//prima stampo gli automi:
	fmt.Println("(")
	for k, v := range p.automi {
		fmt.Printf("%s: %d,%d\n", k, v.x, v.y)
	}
	fmt.Println(")")
	fmt.Println("[")
	p.ostacoli.StampaRettangoli()
	fmt.Println("]")
}

func automa(p piano, x int, y int, n string) {

	punto := punto{x, y}
	//controllo se si conosce già cosa c'è nel relativo punto
	s, esiste := p.puntiConosciuti[punto]
	if esiste {
		switch s {
		case "A":
			p.puntiConosciuti[punto] = "E"
			p.automi[n] = punto
		case "E":
			p.automi[n] = punto
			p.puntiConosciuti[punto] = "A"
		}
		return
	}
	//se non si consoce bisogna per forza cercare
	if !p.ostacoli.ContienePunto(punto) {
		p.automi[n] = punto
		p.puntiConosciuti[punto] = "A"
	} else {
		p.puntiConosciuti[punto] = "O"
	}

	// costo: O(r) con r = numero di rettangoli

}

// dato un nuovo rettangolo aggiunto, aggiunge i punti del perimetro di tale rettangolo alla mappa dei punti conosciuti
func aggiungiPerimetro(p piano, x0 int, y0 int, x1 int, y1 int) {
	//aggiunta dei lati orizzontali
	for i := x0; i <= x1; i++ {
		p.puntiConosciuti[punto{i, y0}] = "O"
		p.puntiConosciuti[punto{i, y1}] = "O"
	}

	//aggiunta dei lati verticali
	for i := y0; i <= y1; i++ {
		p.puntiConosciuti[punto{x0, i}] = "O"
		p.puntiConosciuti[punto{x1, i}] = "O"
	}

	// Costo: O(n+m) con n = x1-x0 e m = y1-y0
}

func ostacolo(p piano, x0 int, y0 int, x1 int, y1 int) {
	p1 := punto{x0, y0}
	p2 := punto{x1, y1}
	rett := rettangolo{p1, p2}
	for _, punto := range p.automi {
		if contiene(rett, punto) {
			return
		}
	}

	p.ostacoli.AddRettangolo(rett)
	//Aggiungo i punti del perimetro del rettangolo poichè un automa, dato un punto di arrivo, può scontrarsi solamente con i punti del perimetro degli ostacoli
	aggiungiPerimetro(p, x0, y0, x1, y1)

	// Costo: O(a + n + m) con a = numero di automi, n = x1-x0 e m = y1-y0

}

func posizioni(p piano, s string) {

	fmt.Println("(")
	for k, v := range p.automi {
		if strings.HasPrefix(k, s) {
			fmt.Printf("%s: %d,%d\n", k, v.x, v.y)
		}
	}
	fmt.Println(")")

	// Costo: O(a * |s|) con a = numero di automi e |s| = lunghezza della stringa s
}

func distanzaPunti(p1 punto, p2 punto) int {
	return int(math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y)))

	// Costo: O(1)
}

//dato un automa e un punto di richiamo, tale automa può muoversi secondo due percorsi unitari, in base a dove è il punto di arrivo. Non può muoversi in tutte e 4 le direzioni
//poichè deve fare un un percorso di distanza minima, e se andasse in un percorso che si allontana dal punto di arrivo, non sarebbe il percorso di distanza minima

func signum(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

// dato un punto di arrivo e un automa in posizione, restituisce le direzioni in cui l'automa può muoversi in un percorso
// minimo. La prima variabile di ritorno indica la direzione sull'asse x, la seconda sull'asse y (e tali valori
// possono essere -1, 0, 1)
func direzioni(puntoAutoma punto, puntoArrivo punto) (int, int) {
	dx := signum(puntoArrivo.x - puntoAutoma.x)
	dy := signum(puntoArrivo.y - puntoAutoma.y)
	return dx, dy
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

// con l'assunzione che il punto di partenza debba avere direzione dx = 1 o 0, percorriamo la tabella sempre da sinistra verso destra
func esistePercorso(p piano, partenza punto, arrivo punto) bool {

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
			if str := p.puntiConosciuti[normalizzaEntrataTabella(partenza, arrivo, diry, i, j)]; str == "O" {
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

func richiamo(p piano, x int, y int, s string) {

	punto := punto{x, y}
	//controllo se si conosce già cosa se c'è un ostacolo in tale punto
	str := p.puntiConosciuti[punto]
	if str == "O" {
		return
	}
	//se non conosco cosa c'è devo vedere se c'è un ostacolo o meno
	if str == "" && p.ostacoli.ContienePunto(punto) {
		p.puntiConosciuti[punto] = "O"
		return
	}

	//se sono arrivato qui vuol dire che non c'è un ostacolo e quindi bisogna sapere gli automi con prefisso s e poi calcolare per ognuno di essi se esiste un percorso libero.
	//Per gli automi che hanno un percorso libero di distanza D allora salvo quelli che possono raggiungere il punto di richiamo. Essi sono quelli con D minima e percorso libero.
	nomiAutomi := automiPrefisso(p, s)
	automiDaSpostare := []string{}
	min := -1
	for _, nome := range nomiAutomi {
		//devo mettere come primo punto quello che deve andare a destra
		p1, p2 := scegliPartenza(p.automi[nome], punto)
		if esistePercorso(p, p1, p2) {
			d := distanzaPunti(p.automi[nome], punto)
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
	for _, nome := range automiDaSpostare {
		p.puntiConosciuti[p.automi[nome]] = "E"
		p.automi[nome] = punto
		p.puntiConosciuti[punto] = "A"
	}
}

// dato un punto e una stringa rappresentante un richiamo, restituisce gli automi con prefisso dato
func automiPrefisso(p piano, s string) []string {
	nomiAutomi := []string{} // ogni volta questa slice viene svuotata e riempita con i nomi degli automi a distanza minima
	for k := range p.automi {
		if strings.HasPrefix(k, s) {
			nomiAutomi = append(nomiAutomi, k)
		}
	}
	return nomiAutomi

	// Costo: O(a * |s|) con a = numero di automi e |s| = lunghezza della stringa s
	// spazio: O(a) nel caso peggiore con a = numero di automi
}

func stato(p piano, x int, y int) {

	punto := punto{x, y}
	s := p.puntiConosciuti[punto]
	if s != "" {
		fmt.Println(s)
		return
	}

	if p.ostacoli.ContienePunto(punto) {
		fmt.Println("O")
		p.puntiConosciuti[punto] = "O"
	} else {
		fmt.Println("E")
		p.puntiConosciuti[punto] = "E"
	}
}
