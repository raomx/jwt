package jwt

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	setSecret("TestVerify")
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1NzkxODc0OTcsImlhdCI6MTU3OTE4NzQ5NywiZXhwIjoxNTc5MTk3MDAwLCJpc3MiOiJhdXRoLmV4YW1wbGUuY29tIn0.Bm6P3FogmZ0KClttChN_TuKes6_yrgo_FylkGNlW404"
	_, err := Verify(token)
	if err == nil {
		t.Error("Verify err")
	}
	var cusClaims = Claims{}
	token = cusClaims.Sign()
	_, err = Verify(token)
	if err != nil {
		t.Error("Verify err")
	}
}

func TestCheckExpires(t *testing.T) {
	var claims = Claims{}
	now := time.Now()
	claims[notBeforeAt] = now.Add(-10 * time.Minute).Unix()
	err := claims.checkExpires(now.Unix(), notBeforeAt)
	if err != nil {
		t.Error("Check NotBeforeAt err")
	}
	claims[notBeforeAt] = now.Add(10 * time.Minute).Unix()
	err = claims.checkExpires(now.Unix(), notBeforeAt)
	if err == nil {
		t.Error("Check NotBeforeAt err")
	}

	claims[expiresAt] = now.Add(-10 * time.Minute).Unix()
	err = claims.checkExpires(now.Unix(), expiresAt)
	if err == nil {
		t.Error("Check ExpiresAt err")
	}
	claims[expiresAt] = now.Add(10 * time.Minute).Unix()
	err = claims.checkExpires(now.Unix(), expiresAt)
	if err != nil {
		t.Error("Check ExpiresAt err")
	}
}

func TestVerifyHead(t *testing.T) {
	var headToke = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	var err = verifyHeader(headToke)
	if err != nil {
		t.Error("VerifyHead err")
	}
	headToke = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9"
	err = verifyHeader(headToke)
	if err == nil {
		t.Error("VerifyHead err")
	}
}

func TestVerify(t *testing.T) {
	setSecret("TestVerify")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.8G4jyYZbsg1xtuUzD0JNSiHfIPT0CziOw1DRCTszVDE"
	wanted := true
	if verifyToken(token) != wanted {
		t.Error("TestVerify failed")
	}
}

func TestGetToken(t *testing.T) {
	setSecret("TestVerify")
	var claims = Claims{}
	claims["jti"] = 123
	claims["sub"] = "test"
	claims["role"] = "user"
	claims["name"] = "rao"

	if !verifyToken(claims.Sign()) {
		t.Error("Sign failed")
	}
}
