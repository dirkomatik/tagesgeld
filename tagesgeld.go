package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("\nMonatliche Tagesgeld-Berechnung\n")

	// liest die werte für die Berechnung ein
	kapital, err := strconv.ParseFloat(readText("Bitte Kapital eingeben: "),64)
	zinssatz, err := strconv.ParseFloat(readText("Bitte Zinsatz eigeben: "),64)
	monate, err := strconv.ParseInt(readText("Bitte Laufzeit in Monate angeben: "),10,64)
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("Kapital: %.0f€, Zinsatz: %.2f%% p.a.\n", kapital, zinssatz)

	// Berechnet die Zinssätze für 1 Jahr
	ergebnis := CalcTagesgeld(kapital, zinssatz, monate)

	fmt.Println("Ertrag: ", ergebnis)
}

func readText(prompt string) string {
	// liest Text von der Kommandozeile ein
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	// list bis zum Zeilenumbruch (der auch mitgelesen wird)
	text,_ := reader.ReadString('\n')

	// entfernt den Zeilenumbruch
	text = strings.TrimRight(text,"\n")

	return text
}

func CalcTagesgeld(kapital float64, zinssatz float64, monate int64) float64 {

	// Zinsatz in Kommaform (5% == 0.05)
	zinssatz = zinssatz / 100.0
	
	// Monat in Kommaform
	monat := 30.0 / 360.0

	zinsertrag := kapital

	// rechnet für jeden Monat den Zinertrag hinzu
	for i := int64(1); i <= monate; i++ {	
		zinsertrag += monat * zinsertrag * zinssatz
	}

	return zinsertrag
}