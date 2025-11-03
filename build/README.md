# Ausführbare Programme mit Go erzeugen

Hier sind die gängigsten Varianten, mit Go ein **ausführbares Programm** zu erzeugen – jeweils kurz erklärt und mit einem Mini-Beispiel.

## 1) `go run` (zum schnellen Ausführen, ohne dauerhaftes Binary)
Baut im Hintergrund und startet sofort – praktisch für Skripte, Demos und Tests, aber legt standardmäßig keine Datei ab.
```bash
go run ./cmd/meinprog
```

## 2) `go build` (lokales Binary im Projekt)
Baut ein ausführbares Programm aus einem `package main`; mit `-o` benennst du die Ausgabe. Ideal für lokale Builds oder CI-Artefakte.
```bash
go build -o meinprog ./cmd/meinprog
```

## 3) `go install` (Binary in $GOBIN / $GOPATH/bin)
Baut und legt das Binary in deinen Bin-Pfad; mit Modulen auch versionierbar (`@vX.Y.Z` oder `@latest`). Gut, um Tools systemweit verfügbar zu machen.
```bash
go install github.com/user/tool/cmd/tool@latest
```

## 4) Cross-Kompilierung (für andere OS/Architekturen)
Über `GOOS`/`GOARCH` erzeugst du Binaries für z. B. Linux, Windows, macOS oder ARM – ohne fremde Toolchains (meist ohne cgo).
```bash
GOOS=linux GOARCH=arm64 go build -o meinprog-linux-arm64 ./cmd/meinprog
```

## 5) Statisch vs. dynamisch linken (cgo)
Reine Go-Programme sind i. d. R. statisch und portable. Sobald cgo gebraucht wird (z. B. SQLite/CGO), entstehen dynamische Abhängigkeiten; mit `CGO_ENABLED=0` erzwingst du statische Builds (falls kompatibel).
```bash
CGO_ENABLED=0 go build -o meinprog-static .
```

## 6) Build-Modi (`-buildmode`) für spezielle Ausgabeartefakte
### **exe/pie**
Normales (Position-Independent) Executable; `pie` verbessert Address Space Layout Randomization.
```bash
go build -buildmode=pie -o meinprog .
```
### **c-archive**
Erstellt ein C-statisches Archiv (`.a`) plus Header – von C aus einbindbar (kein eigenständiges Exe, aber „ausführbar“ im größeren System).
```bash
go build -buildmode=c-archive -o libmeinprog.a .
```
### **c-shared**
Shared Library (`.so`/`.dll`) zum Laden durch andere Programme.
```bash
go build -buildmode=c-shared -o libmeinprog.so .
```
### **plugin**
Go-Plugin (`.so`) zum dynamischen Laden via `plugin`-Paket (Linux/macOS).
```bash
go build -buildmode=plugin -o filter.so ./plugin/filter
```

## 7) WebAssembly (im Browser oder Wasm-Runtime ausführbar)
Baut ein `.wasm`, das im Browser/Runtime läuft; praktisch für UIs oder Sandbox-Ausführung, auch wenn es kein „native“ OS-Exe ist.
```bash
GOOS=js GOARCH=wasm go build -o app.wasm ./web
```

## 8) Race-Detektor-Builds (debug-fähige Executables)
Erzeugt ein instrumentiertes Binary zur Laufzeit-Erkennung von Datenrennen – größer/langsamer, aber sehr nützlich in Tests/Debugging.
```bash
go build -race -o meinprog-race .
```

## 9) Größen-/Info-optimierte Builds (Release-Tuning)
Mit Linker-Flags entfernst du Debug-Symbole, setzt Variablen oder streifst Pfade – gut für kleine, reproduzierbare Binaries.
```bash
go build -ldflags="-s -w -X main.version=v1.2.3" -trimpath -o meinprog .
```

## 10) Container-Executables (als OCI-Image lauffähig)
Du baust ein statisches Go-Binary und packst es in ein minimales Image (z. B. `scratch`). Aus Sicht der Plattform ist das Image das „ausführbare Artefakt“.
```Dockerfile
# build stage
FROM golang:1.23 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/app ./cmd/app

# run stage
FROM scratch
COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
```

## 11) Alternativer Compiler: **TinyGo** (MCUs & sehr kleine Binaries)
Für Embedded, WASM und extrem kleine Ausgaben. Nicht alle Standardbibliotheken/Features werden unterstützt, aber ideal, wenn Ressourcen knapp sind.
```bash
tinygo build -o app.wasm -target wasm ./cmd/app
```

---

**Kurzfazit:** Für die meisten Fälle nimmst du `go build`/`go install` (ggf. mit Cross-Compile-Env-Vars und Release-Flags). Spezialfälle decken `-buildmode` (Libs/Plugins), WebAssembly, Container-Images und TinyGo ab.

