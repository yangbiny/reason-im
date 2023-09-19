package web

import (
	"context"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	token, err := GenerateJwtToken(context.Background(), "username", 123)
	if err != nil {
		t.Fatal(err)
	}
	jwtToken, _ := ParseJwtToken(token)
	println(jwtToken)
	t.Log(token)
}

func TestParse(t *testing.T) {

}
