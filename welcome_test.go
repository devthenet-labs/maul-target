package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWelcomeHandlerNamesTeam(t *testing.T) {
	rec := httptest.NewRecorder()
	welcomeHandler(rec, httptest.NewRequest("GET", "/welcome?team=platform", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "platform") {
		t.Fatalf("response does not name the team: %q", rec.Body.String())
	}
}

func TestWelcomeHandlerEscapesHTML(t *testing.T) {
	rec := httptest.NewRecorder()
	welcomeHandler(rec, httptest.NewRequest("GET", "/welcome?team=%3Cscript%3Ealert(1)%3C%2Fscript%3E", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "<script>") {
		t.Fatalf("response contains unescaped script tag: %q", rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "&lt;script&gt;") {
		t.Fatalf("response does not contain escaped team name: %q", rec.Body.String())
	}
}
