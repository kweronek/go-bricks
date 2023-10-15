# go-bricks
Dieses Repository gibt ein Set von verschiedenen kleinen Go-Programmen. An diesen lassen sich leicht die Syntax und Funktionsweise von go-Sprachelementen erkennen. Darüber hinaus können die Beispiel as Templates eingesetzt werden.
Die Beispiele sind untergliedert in:
* Elementary
* Foundation
* Advanced

## Elementary
* helloWorld
* strings
* arrays    
* ifelse     
* loops          
* switchcase 
* timestamps

## Foundation
* func
* struct
* methods 
* pointer 
* ranges  
* maps    
* files   

## Advanced
* interfaces 
* slices
* channel
* dirfiles
* iota       
* likeEnum

## Error handling
Eines der herausragenden Merkmale von Go ist sein Ansatz zur Fehlerbehandlung, der sich von vielen anderen Programmiersprachen unterscheidet. Statt auf Ausnahmen (Exceptions) zu setzen, wie viele andere Sprachen, verwendet Go einen einzigartigen Ansatz, der sich auf Rückgabewerte konzentriert.
**Einführung in Fehlerbehandlung in Golang**

**1. Einleitung**

Go, oft als "Golang" bezeichnet, ist für seine Performance, Einfachheit und effiziente Konkurrenzkontrolle bekannt. Ein weiterer bemerkenswerter Aspekt von Go ist sein Ansatz zur Fehlerbehandlung. Im Gegensatz zu vielen anderen Programmiersprachen, die auf Ausnahmemechanismen (Exceptions) setzen, verfolgt Go einen anderen Weg: explizite Fehler als Rückgabewerte. Dieser Ansatz fördert die Schreibweise von klarem und vorhersagbarem Code.

**2. Go's Konzeption von Fehlern: Mehrere Rückgabewerte**

Go unterstützt die Rückgabe mehrerer Werte aus Funktionen. In der Praxis bedeutet dies, dass viele Funktionen sowohl ihr eigentliches Ergebnis als auch einen möglichen Fehler als Rückgabewerte haben:

```go
func Dividieren(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("Durch Null teilen")
    }
    return a / b, nil
}
```

Ein solcher Code stellt sicher, dass Fehler nicht stillschweigend ignoriert werden, sondern explizit behandelt werden müssen.

**3. Der `error` Typ in Go**

Im Herzen von Go's Fehlerbehandlungsphilosophie steht das `error` Interface:

```go
type error interface {
    Error() string
}
```

Jeder benutzerdefinierte Typ, der die `Error()` Methode implementiert, kann als Fehler verwendet werden. Das Standardpaket `errors` bietet einfache Funktionen zur Erzeugung von Fehlerobjekten.

**4. Fehlerüberprüfung: Das Kernprinzip**

Die grundlegende Idee der Fehlerbehandlung in Go ist die explizite Überprüfung:

```go
ergebnis, err := Dividieren(10, 2)
if err != nil {
    fmt.Println("Fehler:", err)
    return
}
fmt.Println("Ergebnis:", ergebnis)
```

Hier wird zuerst überprüft, ob `err` ungleich `nil` ist. Ist dies der Fall, wurde ein Fehler zurückgegeben, und die entsprechende Fehlerbehandlung wird ausgeführt.

**5. Benutzerdefinierte Fehler**

Einfache Fehlermeldungen sind oft ausreichend, aber manchmal benötigt man detailliertere Informationen:

```go
type MeinFehler struct {
    Nachricht string
    Code      int
}

func (e MeinFehler) Error() string {
    return fmt.Sprintf("Code: %d, Nachricht: %s", e.Code, e.Nachricht)
}
```

Durch solche benutzerdefinierten Fehlerstrukturen kann man zusätzliche Kontextinformationen bereitstellen, was insbesondere bei großen und komplexen Systemen nützlich ist.

**6. Best Practices und Tipps**

- **Explizite Überprüfung**: Go ermutigt Entwickler, jeden möglichen Fehler sofort nach dem Aufruf der Funktion zu überprüfen.

- **"Sentinel" Fehler**: Das sind vordefinierte Fehler, die für den Vergleich verwendet werden können. Zum Beispiel: `io.EOF` signalisiert, dass das Ende einer Datei erreicht wurde.

- **Wrapper-Funktionen**: Das `fmt.Errorf` erlaubt es, Fehlermeldungen zu formatieren und dabei einen anderen Fehler zu "umwickeln", wodurch zusätzlicher Kontext hinzugefügt werden kann.

- **Vermieden von "panic"**: Während Go die Möglichkeit bietet, mit `panic` und `recover` einen ähnlichen Mechanismus wie Exceptions in anderen Sprachen zu nutzen, sollte dies in regulären Anwendungen vermieden werden. Panics sind für nicht-behebbare Laufzeitfehler gedacht, nicht für normale Fehlerbehandlung.

**7. Schlussfolgerung**

Go's Ansatz zur Fehlerbehandlung unterscheidet sich deutlich von dem vieler anderer Sprachen. Durch die Förderung expliziter Fehlerüberprüfungen hilft Go Entwicklern, robusten, wartbaren und verständlichen Code zu schreiben. Fehler werden nicht leichtfertig übersehen, und Entwickler sind stets im Bilde über mögliche Problemstellen ihrer Anwendungen. Dieser Ansatz mag zunächst ungewöhnlich erscheinen, hat sich aber in der Praxis als effektiv und effizient erwiesen.
## Testing 
tbd

## System-related
tbd
