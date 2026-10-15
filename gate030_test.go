package main

import "testing"

func TestGate030KeyMeetsMinimumSize(t *testing.T) {
	key, err := gate030Key()
	if err != nil {
		t.Fatalf("gate030Key() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("gate030Key() produced a %d-bit key, want >= 2048", bits)
	}
}
