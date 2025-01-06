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
		if esistePercorso(p, a, b, comandi[3]) {
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

// Questa funzione aggiunge i vicini da considerare in un percorso di distanza minima. Se dx != 0, allora l'automa può muoversi in orizzontale, se dy != 0, allora l'automa può muoversi in verticale
func aggiungiAdiacenti(dx int, puntoCorrente punto, p piano, coda *MyCoda, distanzaPunto int, dy int) {
	if dx != 0 {
		puntoAdiacente := punto{puntoCorrente.x + dx, puntoCorrente.y}
		s := p.puntiConosciuti[puntoAdiacente]
		if s != "O" {
			coda.enqueue(puntoAdiacente, distanzaPunto+1)
		}
	}
	if dy != 0 {
		puntoAdiacente := punto{puntoCorrente.x, puntoCorrente.y + dy}
		s := p.puntiConosciuti[puntoAdiacente]
		if s != "O" {
			coda.enqueue(puntoAdiacente, distanzaPunto+1)
		}
	}
}

// dato un punto di partenza, un punto di arrivo, un automa in posizione, le direzioni in cui l'automa può muoversi e la distanza minima tra il punto di partenza e di arrivo,
// restituisce true se esiste un percorso libero di distanza minima tra i due punti. Un percorso è libero se non incontra ostacoli. Inoltre si possono incontrare ostacoli
// che rappresentano solo punti del perimetro di un rettangolo, i quali sono stati aggiunti alla mappa dei punti conosciuti e quindi accessibili in tempo costante
func bfs(p piano, puntoPartenza punto, puntoArrivo punto, dx int, dy int, distanzaD int) bool {

	coda := &MyCoda{}
	coda.enqueue(puntoPartenza, 0)
	//visitati := make(map[Punto]bool)
	for !coda.isEmpty() {
		punto, distanzaPunto := coda.dequeue()
		/*
			//se il punto è già stato visitato, allora non lo visito di nuovo
			if visitati[punto] {
				continue
			}
			visitati[punto] = true
		*/
		if distanzaPunto == distanzaD && punto == puntoArrivo {
			return true
		}
		//se la distanza di tale punto è uguale a quella di arrivo e non è il punto di arrivo, allora non visito i punti adiacenti
		if distanzaPunto == distanzaD {
			continue
		}
		//aggiungo i punti adiacenti
		aggiungiAdiacenti(dx, punto, p, coda, distanzaPunto, dy)

	}
	return false

}

// dato un punto di arrivo (x, y) e un automa in posizione (x1, y1), restituisce true se esiste un percorso libero di distanza minima tra i due punti. Un percorso è libero se non incontra
// ostacoli
func esistePercorso(p piano, x int, y int, s string) bool {

	puntoAutoma, esiste := p.automi[s]
	if !esiste {
		return false
	}
	puntoArrivo := punto{x, y}
	distanzaD := distanzaPunti(puntoAutoma, puntoArrivo)
	if distanzaD == 0 {
		return true
	}
	/*
		if p.ostacoli.ContienePunto(puntoArrivo) {
			p.puntiConosciuti[puntoArrivo] = "O"
			return false
		}
	*/

	//calcolo la direzione in cui l'automa deve muoversi
	dx, dy := direzioni(puntoAutoma, puntoArrivo)

	//creo la coda per la visita in ampiezza. Ogni elemento della coda è un punto e la distanza D da tale punto a puntoAutoma. Mi fermo ad aggiungere elementi nella coda quando
	//raggiungo punti con distanza distanzaD. Se arrivo a puntoArrivo, allora esiste un percorso libero, se no non esiste perchè ho fatto tutti i percorsi possibili di distanza minima e non ho trovato
	//puntoArrivo

	return bfs(p, puntoAutoma, puntoArrivo, dx, dy, distanzaD)

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
		if esistePercorso(p, x, y, nome) {
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
