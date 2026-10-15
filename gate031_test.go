package main

import "testing"

func TestGate031KeyMeetsMinimumSize(t *testing.T) {
	key, err := gate031Key()
	if err != nil {
		t.Fatalf("gate031Key() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("gate031Key() produced a %d-bit key, want >= 2048", bits)
	}
}
