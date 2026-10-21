package main

import (
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestWordcountHandlerCountsWords(t *testing.T) {
	rec := httptest.NewRecorder()
	wordcountHandler(rec, httptest.NewRequest("GET", "/wordcount?text=one+two+three", nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "3 words") {
		t.Fatalf("response does not report three words: %q", rec.Body.String())
	}
}

func TestWordcountHandlerDoesNotExecuteShellMetacharacters(t *testing.T) {
	rec := httptest.NewRecorder()
	wordcountHandler(rec, httptest.NewRequest("GET", "/wordcount?text="+url.QueryEscape("foo; touch /tmp/maul-pwned; echo bar"), nil))
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if _, err := os.Stat("/tmp/maul-pwned"); err == nil {
		os.Remove("/tmp/maul-pwned")
		t.Fatalf("shell metacharacters were executed: /tmp/maul-pwned was created")
	}
	if !strings.Contains(rec.Body.String(), "5 words") {
		t.Fatalf("response does not report the literal word count: %q", rec.Body.String())
	}
}
