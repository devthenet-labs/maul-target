package main

import "testing"

func TestGate032KeyMeetsMinimumSize(t *testing.T) {
	key, err := gate032Key()
	if err != nil {
		t.Fatalf("gate032Key() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("gate032Key() produced a %d-bit key, want >= 2048", bits)
	}
}
