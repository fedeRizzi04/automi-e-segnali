package main

import (
	"fmt"
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

	p := newPiano()

	p.automi["1"] = punto{1, 1}
	p.automi["11"] = punto{2, 2}
	p.automi["111"] = punto{3, 3}
	p.automi["1111"] = punto{4, 4}
	p.automi["11111"] = punto{5, 5}
	p.automi["111111"] = punto{6, 6}
	p.automi["1111111"] = punto{7, 7}
	p.automi["11111111"] = punto{8, 8}
	p.automi["111111111"] = punto{9, 9}
	p.automi["1111111111"] = punto{10, 10}
	p.automi["11111111111"] = punto{11, 11}
	p.automi["111111111111"] = punto{12, 12}
	p.automi["1111111111111"] = punto{13, 13}
	p.automi["11111111111111"] = punto{14, 14}
	p.automi["111111111111111"] = punto{15, 15}
	p.automi["1111111111111111"] = punto{16, 16}
	p.automi["11111111111111111"] = punto{17, 17}
	p.automi["111111111111111111"] = punto{18, 18}
	p.automi["1111111111111111111"] = punto{19, 19}
	p.automi["11111111111111111111"] = punto{20, 20}
	p.automi["111111111111111111111"] = punto{21, 21}
	p.automi["1111111111111111111111"] = punto{22, 22}
	p.automi["11111111111111111111111"] = punto{23, 23}
	p.automi["111111111111111111111111"] = punto{24, 24}
	p.automi["1111111111111111111111111"] = punto{25, 25}
	p.automi["11111111111111111111111111"] = punto{26, 26}
	p.automi["111111111111111111111111111"] = punto{27, 27}
	p.automi["1111111111111111111111111111"] = punto{28, 28}
	p.automi["11111111111111111111111111111"] = punto{29, 29}
	p.automi["111111111111111111111111111111"] = punto{30, 30}
	p.automi["1111111111111111111111111111111"] = punto{31, 31}
	p.automi["11111111111111111111111111111111"] = punto{32, 32}
	p.automi["111111111111111111111111111111111"] = punto{33, 33}
	p.automi["1111111111111111111111111111111111"] = punto{34, 34}
	p.automi["11111111111111111111111111111111111"] = punto{35, 35}
	p.automi["0111"] = punto{5, 5}
	p.automi["00111"] = punto{6, 6}

	sl := automiPrefisso(p, "1")
	fmt.Println(sl)

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

// FUNZIONI PER SPECIFICHE DI PROGETTAZIONE

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
	fmt.Println(len(nomiAutomi), cap(nomiAutomi))
	rangeAutomiPrefisso(p, s, func(k string, p punto) {
		nomiAutomi = append(nomiAutomi, k)
	})
	return nomiAutomi
}
