# Automi e segnali

Questo repository contiene il progetto realizzato per l'insegnamento di "Algoritmi e strutture dati" della Laurea triennale in Informatica all'Università degli Studi di Milano. L'anno accademico è il 2024/2025.

In questo progetto viene rappresentato un piano cartesiano potenzialmente infinito sul quale possono essere inseriti automi e ostacoli rettangolari. Ogni automa è identificato da un nome univoco sull'alfabeto `{0, 1}` (oltre che dalle coordinate $x$ ed $y$ che descrivono il punto su cui giace l'automa), mentre ogni ostacolo è definito dal proprio vertice in basso a sinistra e dal proprio vertice in alto a destra.

## Notazione e percorsi

Sia $\eta$ il nome di un automa, con $P(\eta)$ si indica il punto del piano sul quale esso si trova. Dati due punti $P=(x_0,y_0)$ e $Q=(x_1,y_1)$, la loro distanza di Manhattan è definita come:

$$
D(P,Q)=|x_1-x_0|+|y_1-y_0|
$$

Un passo unitario collega due punti adiacenti in direzione orizzontale o verticale e modifica quindi una sola coordinata di $1$ o $-1$. Un percorso è una sequenza di passi unitari e la sua lunghezza corrisponde al numero di passi che lo compongono. Per andare da $P$ a $Q$ sono necessari almeno $|x_1-x_0|$ passi orizzontali e $|y_1-y_0|$ passi verticali: un percorso ha quindi lunghezza minima quando la sua lunghezza è esattamente $D(P,Q)$.

Un percorso è libero se nessuno dei punti attraversati è occupato da un ostacolo. Non sono ammessi percorsi più lunghi che aggirano un ostacolo: l'operazione richiesta deve stabilire se esiste un percorso libero di lunghezza esattamente pari alla distanza di Manhattan. Ai fini dell'operazione tale lunghezza deve inoltre essere maggiore di zero; se $P(\eta)=Q$, quindi, il risultato è negativo.

Nella descrizione delle operazioni, $R(x_0,y_0,x_1,y_1)$ indica un ostacolo rettangolare e $s$ indica un prefisso sull'alfabeto `{0, 1}`.

## Regole del piano

Gli automi e gli ostacoli vengono inseriti sul piano rispettando le seguenti regole:

- più automi possono trovarsi sullo stesso punto, purché abbiano nomi differenti;
- un automa non può essere inserito o spostato su un punto occupato da un ostacolo;
- un ostacolo può essere aggiunto solamente se la propria area non contiene automi;
- gli ostacoli occupano anche i punti appartenenti al proprio perimetro e, una volta inseriti, non possono essere rimossi.

Gli automi non costituiscono un ostacolo per il movimento. Se viene inserito un automa con un nome già presente, esso viene riposizionato nel nuovo punto, a condizione che questo non sia occupato da un ostacolo.

## Operazioni principali

Oltre all'inserimento di automi e ostacoli, è possibile conoscere lo stato di un punto, stampare tutte le entità presenti nel piano e cercare gli automi il cui nome ha un determinato prefisso.

L'operazione di esistenza di un percorso, dati un automa $\eta$ e un punto di arrivo $Q$, stabilisce se esiste un percorso libero da $P(\eta)$ a $Q$ di lunghezza $D(P(\eta),Q)$. Il risultato è negativo se l'automa non esiste, se un ostacolo impedisce tutti i percorsi di lunghezza minima oppure se l'automa si trova già sul punto di arrivo.

L'operazione di richiamo, dati un punto $Q$ e un prefisso $s$, considera tutti gli automi il cui nome inizia con $s$. Tra gli automi per cui esiste un percorso libero verso $Q$ vengono spostati tutti e soli quelli per cui $D(P(\eta),Q)$ è minima. Se il punto di richiamo è occupato da un ostacolo non viene spostato alcun automa.

La relazione completa, nella quale vengono descritte le scelte implementative, gli algoritmi utilizzati e l'analisi dei costi, è resa disponibile in formato PDF nelle [release del repository](../../releases). Il sorgente Latex è disponibile all'interno del repository (_relazione.tex_). All'interno della relazione vengono approfondite le operazioni spiegate in questo paragrafo.

## Esecuzione

Il programma richiede Go e legge i comandi dallo standard input. Può essere eseguito direttamente con:

```bash
go run main.go
```

In alternativa è possibile passare un file contenente una sequenza di comandi:

```bash
go run main.go < tests/input2.txt
```

Il primo comando deve essere `c`, che crea un nuovo piano. I comandi disponibili sono:

- `c`: crea un nuovo piano vuoto;
- `a x y nome`: inserisce un automa nel punto $(x, y)$ oppure aggiorna la posizione dell'automa che ha lo stesso nome;
- `o x0 y0 x1 y1`: inserisce un ostacolo rettangolare con vertice in basso a sinistra $(x0, y0)$ e vertice in alto a destra $(x1, y1)$, purché non contenga automi;
- `s x y`: stampa `A` se nel punto è presente almeno un automa, `O` se è presente un ostacolo ed `E` se il punto è vuoto;
- `S`: stampa tutti gli automi e tutti gli ostacoli presenti nel piano;
- `p prefisso`: stampa posizione e nome degli automi il cui nome inizia con il prefisso specificato;
- `e x y nome`: stampa `SI` se esiste un percorso libero di lunghezza minima tra l'automa e il punto $(x, y)$, `NO` altrimenti;
- `r x y prefisso`: richiama verso $(x, y)$ gli automi con il prefisso specificato, spostando quelli raggiungibili che hanno distanza minima dal punto;
- `f`: termina il programma.

## Test

La cartella `tests` contiene sei coppie di file. Ogni file `inputX.txt` contiene i comandi da fornire al programma, mentre il corrispondente `expectedX.txt` contiene l'output atteso. Per esempio, il secondo test può essere controllato con:

```bash
go run main.go < tests/input2.txt | diff -u tests/expected2.txt -
```

I test prendono in considerazione piani con entità molto sparse, percorsi su una sola direzione, ostacoli numerosi e diversi casi limite. Sono inoltre presenti esempi relativi sia all'esistenza di un percorso sia al richiamo degli automi.
