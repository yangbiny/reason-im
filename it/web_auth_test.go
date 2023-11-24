package it

import (
	"reason-im/internal/config/web"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	err := web.GenerateJwtToken(nil, "username", 123)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {

}
