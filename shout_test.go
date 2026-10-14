package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShoutHandlerUppercases(t *testing.T) {
	rec := httptest.NewRecorder()
	shoutHandler(rec, httptest.NewRequest("GET", "/shout?msg=hello+there", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "HELLO THERE") {
		t.Fatalf("response does not shout the message: %q", rec.Body.String())
	}
}

func TestShoutHandlerEscapesHTML(t *testing.T) {
	rec := httptest.NewRecorder()
	shoutHandler(rec, httptest.NewRequest("GET", "/shout?msg=%3Cscript%3Ealert(1)%3C%2Fscript%3E", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "<script>") {
		t.Fatalf("response contains unescaped script tag: %q", rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "&lt;SCRIPT&gt;") {
		t.Fatalf("response does not contain escaped message: %q", rec.Body.String())
	}
}
