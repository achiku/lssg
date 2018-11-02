package main

import (
	"os"
	"testing"
)

func TestParseSG(t *testing.T) {
	p := os.Getenv("LSSG_TEST_TF")
	f, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := parseSG(f); err != nil {
		t.Fatal(err)
	}
}
