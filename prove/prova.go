package main

type Punto struct {
	x int
	y int
}

type piano *struct {
	punti []Punto
	m     map[string]int
}

func modifica(p piano) {

	p.m["ciao2"] = 2
	p.m["ciao3"] = 3
	p.m["ciao4"] = 4
	p.m["ciao5"] = 5
	p.m["ciao6"] = 6
	p.m["ciao7"] = 7
	p.m["ciao8"] = 8
	p.m["ciao9"] = 9
	p.m["ciao10"] = 10
	p.m["ciao11"] = 11
	p.m["ciao12"] = 12
	p.m["ciao13"] = 13
	p.m["ciao14"] = 14
	p.m["ciao15"] = 15
	p.m["ciao16"] = 16
	p.m["ciao17"] = 17
	p.m["ciao18"] = 18
	p.m["ciao19"] = 19
	p.m["ciao20"] = 20
	p.m["ciao21"] = 21
	p.m["ciao22"] = 22
	p.m["ciao23"] = 23
	p.m["ciao24"] = 24
	p.m["ciao25"] = 25
	p.m["ciao26"] = 26
	p.m["ciao27"] = 27
	p.m["ciao28"] = 28
	p.m["ciao29"] = 29
	p.m["ciao30"] = 30
	p.m["ciao31"] = 31
	p.m["ciao32"] = 32
	p.m["ciao33"] = 33
	p.m["ciao34"] = 34
	p.m["ciao35"] = 35
	p.m["ciao36"] = 36
	p.m["ciao37"] = 37
	p.m["ciao38"] = 38
	p.m["ciao39"] = 39
	p.m["ciao40"] = 40
	p.m["ciao41"] = 41
	p.m["ciao42"] = 42
	p.m["ciao43"] = 43
	p.m["ciao44"] = 44
	p.m["ciao45"] = 45
	p.m["ciao46"] = 46
	p.m["ciao47"] = 47
	p.m["ciao48"] = 48
	p.m["ciao49"] = 49
	p.m["ciao50"] = 50

	p.m = make(map[string]int, 2)
	p.m["ciao"] = 17

}

func main() {

	var p piano = &struct {
		punti []Punto
		m     map[string]int
	}{
		punti: []Punto{},
		m:     make(map[string]int, 2),
	}

	p.m["ciao"] = 1

	modifica(p)

	for k, v := range p.m {
		println(k, v)
	}

}
