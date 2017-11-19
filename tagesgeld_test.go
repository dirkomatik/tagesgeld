package main

import "testing" 

func TestCalcTagesgeld(t *testing.T) {
	testwert := CalcTagesgeld(5000.0, 5.0, 1)
	sollwert := 5000 + (30.0/360) * 5000 * 0.05

	if testwert != sollwert {
		t.Errorf("Erwartet: %f, gefunden %f", sollwert,testwert)
	}
}