package main

import "testing"

func TestDNSGateKeyMeetsMinimumSize(t *testing.T) {
	key, err := dnsGateKey()
	if err != nil {
		t.Fatalf("dnsGateKey() error = %v", err)
	}
	if bits := key.N.BitLen(); bits < 2048 {
		t.Fatalf("dnsGateKey() produced a %d-bit key, want >= 2048", bits)
	}
}
