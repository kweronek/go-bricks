# Programmierbausteine für Go
Dieses Repository gibt ein Set von verschiedenen kleinen Go-Programmen. An diesen lassen sich leicht die Syntax und Funktionsweise von go-Sprachelementen erkennen. Darüber hinaus können die Beispiel as Templates eingesetzt werden.
Die Beispiele sind untergliedert in:
* Elementary
* Foundation
* Advanced

**1. Elementary**  
* helloWorld
* strings
* arrays    
* ifelse     
* loops          
* switchcase 
* timestamps

**2. Foundation**
* func
* struct
* methods 
* pointer 
* ranges  
* maps    
* files   

**3. Advanced**
* interfaces 
* slices
* channel
* dirfiles
* iota       
* likeEnum

# Error handling
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

# Einführung in das Testing mit Go

**1. Einleitung**

Go, oder oft als "Golang" bezeichnet, ist nicht nur für seine einfache Syntax und effiziente Ausführung bekannt, sondern auch für seine eingebauten Testing-Funktionen. Testing ist ein zentrales Element jeder modernen Softwareentwicklung, und Go bietet den Entwicklern solide Werkzeuge, um den Code zu testen und die Softwarequalität zu sichern. Diese Einführung führt Sie durch die Grundlagen des Testens in Go.

**2. Grundlegende Teststruktur**

Ein Test in Go ist eine Funktion, die in einer speziellen *_test.go Datei lebt. Der Testname beginnt mit `Test` gefolgt vom Namen der Funktion, die getestet wird:

```go
func TestMeineFunktion(t *testing.T) {
    // Ihr Testcode hier
}
```

Der Parameter `t` ist von Typ `*testing.T` und wird für Teststeuerung und -berichterstattung verwendet.

**3. Asserts und Überprüfungen**

Go’s Standardbibliothek für Tests enthält keine traditionellen "Assert"-Funktionen. Stattdessen überprüfen Sie Bedingungen mit normalen Go-Code und nutzen die `t.Error()` oder `t.Fatal()` Methoden für Fehler:

```go
result := MeineFunktion()
if result != erwartet {
    t.Errorf("Erwartet %v, erhalten %v", erwartet, result)
}
```

**4. Tabellengesteuerte Tests**

Eine gängige Praxis in Go ist die Verwendung von tabellengesteuerten Tests, um verschiedene Eingaben und erwartete Ausgaben in einer Tabelle zu definieren:

```go
func TestDividieren(t *testing.T) {
    tests := []struct{
        inputA, inputB, erwartet int
    }{
        {10, 5, 2},
        {12, 4, 3},
        // Weitere Testfälle
    }

    for _, tt := range tests {
        result := Dividieren(tt.inputA, tt.inputB)
        if result != tt.erwartet {
            t.Errorf("Für %v und %v erwartet %v, erhalten %v", tt.inputA, tt.inputB, tt.erwartet, result)
        }
    }
}
```

**5. Mocking und Stubs**

Mocking ist eine Technik, bei der echte Abhängigkeiten durch gefälschte Versionen (Mocks) ersetzt werden. Es gibt mehrere Bibliotheken in Go's Ökosystem, wie `gomock` oder `testify`, die Mocking unterstützen. Diese erlauben es, Verhalten und Erwartungen für externe Abhängigkeiten, wie Datenbankaufrufe, zu definieren.

**6. Benchmarking**

Go bietet auch Unterstützung für Leistungs-Benchmarks. Benchmark-Funktionen leben in *_test.go Dateien, genau wie Tests, beginnen aber mit `Benchmark`:

```go
func BenchmarkMeineFunktion(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MeineFunktion()
    }
}
```

Der Code im Benchmark wird so oft wie angegeben durch den Wert von `b.N` ausgeführt, und Go misst, wie lange es dauert.

**7. Testabdeckung**

Um die Testabdeckung zu überprüfen, kann Go einen Bericht erstellen, der anzeigt, welche Teile des Codes von Tests getroffen wurden. Dies wird über den Befehl `go test -cover` erreicht.

**8. Fazit**

Das Testen in Go ist durch die Sprachdesign-Entscheidungen stark vereinfacht und in den Entwicklungsworkflow integriert. Durch die Nutzung der eingebauten Testing-Werkzeuge, kombiniert mit Best Practices und externen Bibliotheken, können Entwickler robusten, zuverlässigen und effizienten Code schreiben. Das Testen sollte immer ein integraler Bestandteil des Entwicklungsprozesses in jedem Go-Projekt sein.

# System-relevante Befehle
Wenn man mit Go arbeitet, gibt es eine Reihe von systemrelevanten Befehlen, die sowohl für die Entwicklung als auch für die Verwaltung von Go-Anwendungen wichtig sind. Diese Befehle werden über das `go`-Werkzeug ausgeführt.

Hier ist eine Einführung in die systemrelevanten Befehle in Go:

1. **go run**
   - **Verwendung:** Führt Go-Quellcode direkt aus, ohne eine ausführbare Datei zu erstellen.
   - **Beispiel:** `go run main.go`

2. **go build**
   - **Verwendung:** Kompiliert Go-Code in eine ausführbare Binärdatei. Es erstellt eine ausführbare Datei im aktuellen Verzeichnis.
   - **Beispiel:** `go build main.go`

3. **go install**
   - **Verwendung:** Kompiliert und installiert ein Go-Binärprogramm oder eine Bibliothek. Für Binärprogramme werden die Binärdateien im `$GOPATH/bin`-Verzeichnis abgelegt.
   - **Beispiel:** `go install mypackage`

4. **go get**
   - **Verwendung:** Lädt und installiert Pakete aus einer URL zu einem Repository (z.B. GitHub). Es wird auch verwendet, um Abhängigkeiten zu aktualisieren.
   - **Beispiel:** `go get -u github.com/user/repo`

5. **go test**
   - **Verwendung:** Führt Tests für das aktuelle Paket aus. Es identifiziert Testfunktionen (die mit `Test` beginnen) und führt sie aus.
   - **Beispiel:** `go test` oder `go test ./...` (um Tests in allen Unterordnern auszuführen).

6. **go fmt**
   - **Verwendung:** Formatiert Go-Quellcode gemäß den Standardformatierungsrichtlinien von Go. Es ist üblich, dies vor dem Commit von Code in ein Repository auszuführen.
   - **Beispiel:** `go fmt`

7. **go doc**
   - **Verwendung:** Zeigt Dokumentation für Go-Code. Es kann auf Paket-, Funktion- oder Symbolbasis verwendet werden.
   - **Beispiel:** `go doc fmt.Println`

8. **go mod**
   - **Verwendung:** Ein Werkzeug zur Modulverwaltung in Go. Mit `go mod init` können Sie ein neues Modul initialisieren, und mit `go mod tidy` können Sie die Abhängigkeiten aufräumen.
   - **Beispiel:** `go mod init mymodule`

9. **go vet**
   - **Verwendung:** Überprüft den Go-Code auf häufige Fehler. Es ist kein Ersatz für Linting-Werkzeuge, aber es fängt viele gängige Programmierfehler ab.
   - **Beispiel:** `go vet`

10. **go env**
   - **Verwendung:** Zeigt Go-Umgebungsvariablen an.
   - **Beispiel:** `go env`

11. **go version**
   - **Verwendung:** Zeigt die aktuell installierte Go-Version an.
   - **Beispiel:** `go version`

Zusammengefasst bietet das `go`-Werkzeug eine umfassende Sammlung von Befehlen, um sowohl die Entwicklung als auch die Verwaltung von Go-Projekten zu erleichtern. Es ist empfehlenswert, sich mit diesen Befehlen vertraut zu machen, um effizient mit Go arbeiten zu können.
