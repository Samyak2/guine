package main

import (
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestQuine(t *testing.T) {
	out, err := exec.Command("go", "run", "quine.go").Output()
	if err != nil {
		t.Errorf("Error in executing go run: %v", err)
	}

	actualProg, err := ioutil.ReadFile("quine.go")
	if err != nil {
		t.Errorf("Error in reading source file: %v", err)
	}

	outS := string(out)
	progS := string(actualProg)

	if len(outS) != len(progS) {
		t.Errorf("Length mistmatch: Actual program is of length %v while output is of length %v", len(actualProg), len(out))
	}

	for i := range out {
		if outS[i] != progS[i] {
			t.Errorf("Mismatch: at index %d, output is %c while actual program is %c", i, outS[i], progS[i])
		}
	}
}
