package main

import "testing"

func TestSubscriptionGate2KeyMeetsMinimumSize(t *testing.T) {
	key, err := subscriptionGate2Key()
	if err != nil {
		t.Fatalf("subscriptionGate2Key() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("subscriptionGate2Key() produced a %d-bit key, want >= 2048", bits)
	}
}
