package main

// Interfaces

import "fmt"

type Polygons interface {
	Perimeter()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("The perimeter of this pentagon is : ",5 * p)
}

type Hexagon int

func (h Hexagon) Perimeter() {
	fmt.Println("The perimeter of this hexagon is :",6 * h)
}

func main() {
	fmt.Println("Demo for interfaces:")
	var myPolygon Polygons

	myPolygon = Pentagon(50)
	myPolygon.Perimeter() // 250

	myPolygon = Hexagon(50)
	myPolygon.Perimeter() // 300
}

// Definition: Interface:
// Ein Interface besteht aus einer Liste von Methoden, die ein Typ besitzen muss, um das Interface zu erfüllen (“implementieren”).
//
// Ob ein Typ ein Interface erfüllt, wird nicht explizit angegeben, sondern
// ergibt sich implizit aus der Definition des Typs und seiner Methoden.
//
// Das Interface wird selber als benutzerdefinierter Typ definiert.
// Es kann also Variablen vom Typ des Interfaces geben. Auf diese Weise wird Polymorphie in Go realisiert.
//
// Ob ein Typ ein bestimmtes Interface erfüllt, kann zur Laufzeit geprüft werden.
// Von dieser Möglichkeit macht die Go Standardbibliothek häufig Gebrauch.
//
// Es ist eine Konvention, dass ein Interface mit nur einer Methode den Namen der Methode mit angehängtem “-er” bekommt.
// Ein Interface, dessen einzige Funktion String() heißt, erhält also den Namen Stringer.
// Im Beispiel müsste das Interface daher eigentlich Perimeterer heißen
//
// Das leere Interface interface{} wird von allen Typen erfüllt.
//
// Einer Variable vom Typ interface{} können also Variablen jedes beliebigen anderen Typs zugewiesen werden.
// So lassen sich Funktionen definieren, die mit beliebigen Argumenten umgehen können.
// Bekanntestes Beispiel dafür sie die Ausgabefunktionen wie fmt.Println(),
// der man Variablen und Konstanten aller Typen übergeben kann.
//
// Mittels der Reflection API prüft die Funktion, welchen konkreten Typ ein Argument hat, und verhält sich entsprechend.
// Das geschieht bei Bedarf auch rekursiv, z.B. bei Verbundtypen oder Arrays.
