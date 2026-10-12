package main

import "testing"

func TestSubscriptionGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := subscriptionGateKey()
	if err != nil {
		t.Fatalf("subscriptionGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("subscriptionGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
