package it

import (
	"context"
	"reason-im/internal/utils/token"
	"testing"
)

func TestGenerateToken(t *testing.T) {

	claims := map[string]string{}
	claims["x"] = "y"
	applyToken, err := token.ApplyToken(context.Background(), 1, claims)
	if err != nil {
		t.Fatal(err)
	}

	parseToken, err := token.ParseToken(context.Background(), *applyToken)
	if err != nil {
		t.Fatal(err)
	}
	println(parseToken)
}
