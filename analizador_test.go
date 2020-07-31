package analizador

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPrintEstadistica(t *testing.T) {
	salidaAnt := os.Stdout
	defer func() {
		os.Stdout = salidaAnt
	}()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	go func() {
		PrintEstadistica("Ángelo Agustín")
		w.Close()
		os.Stdout = salidaAnt
	}()
	stdoutb, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	stdout := string(stdoutb)

	for _, contiene := range []string {
		"2 mayúsculas",
		"11 minúsculas",
		"6 vocales",
		"7 consonantes",
		"1 carácteres desconocidos",
		"A : 2", "N : 2", "G : 2", "E : 1", "L : 1",
		"O : 1", "U : 1", "S : 1", "T : 1", "I : 1",
	} {
		if !strings.Contains(stdout, contiene) {
			t.Errorf("No encuentro %q en: %s", contiene, stdout)
		}
	}
}
