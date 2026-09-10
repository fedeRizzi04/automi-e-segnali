# Automi e segnali

In questo progetto viene rappresentato un piano cartesiano potenzialmente infinito sul quale possono essere inseriti automi e ostacoli rettangolari. Ogni automa è identificato da un nome univoco sull'alfabeto `{0, 1}` e occupa un punto del piano, mentre un ostacolo occupa tutti i punti compresi tra il proprio vertice in basso a sinistra e quello in alto a destra.

Le operazioni principali consistono nel determinare se un automa può raggiungere un punto attraverso un percorso libero di lunghezza minima e nell'eseguire un richiamo. In quest'ultimo caso vengono considerati gli automi il cui nome ha un determinato prefisso; tra quelli per cui esiste un percorso libero vengono spostati sul punto di richiamo tutti e soli gli automi che si trovano alla distanza minima.

## Approccio utilizzato

La distanza tra due punti è la distanza di Manhattan. Di conseguenza, un percorso di lunghezza minima può muoversi solamente nelle direzioni che avvicinano il punto di partenza a quello di arrivo. I punti candidati a far parte del percorso formano quindi un rettangolo, rappresentato implicitamente mediante una matrice di programmazione dinamica.

Ogni entrata della matrice indica se il punto corrispondente è raggiungibile tramite un percorso libero di lunghezza minima. La matrice viene popolata esaminando per ogni punto i due possibili predecessori; i punti appartenenti al perimetro di un ostacolo vengono invece marcati come non raggiungibili. Sia $r$ il numero di righe e $c$ il numero di colonne del rettangolo compreso tra i due punti, il costo della verifica è $O(r \cdot c)$ sia in tempo che in spazio.

Gli automi sono memorizzati in una mappa che associa ogni nome alla sua posizione. Gli ostacoli sono memorizzati in una lista concatenata, mentre una seconda mappa contiene le informazioni già conosciute sui punti del piano. In particolare vengono memorizzati i punti occupati dagli automi, i punti vuoti già esaminati e il perimetro degli ostacoli. Questo permette di verificare in tempo costante se un punto esaminato durante la ricerca del percorso appartiene al perimetro di un ostacolo.

Una descrizione più approfondita delle strutture dati, degli algoritmi e dei relativi costi è presente in [`relazione.tex`](relazione.tex).

## Esecuzione

Il programma richiede Go e legge i comandi dallo standard input. Può essere eseguito direttamente con:

```bash
go run main.go
```

In alternativa è possibile passare un file contenente una sequenza di comandi:

```bash
go run main.go < test/input2.txt
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

La cartella `test` contiene sei coppie di file. Ogni file `inputX.txt` contiene i comandi da fornire al programma, mentre il corrispondente `expectedX.txt` contiene l'output atteso. Per esempio, il secondo test può essere controllato con:

```bash
go run main.go < test/input2.txt | diff -u test/expected2.txt -
```

I test prendono in considerazione piani con entità molto sparse, percorsi su una sola direzione, ostacoli numerosi e casi limite della matrice di programmazione dinamica. Sono inoltre presenti esempi relativi sia all'esistenza di un percorso sia al richiamo degli automi.
