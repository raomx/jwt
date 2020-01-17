// Copyright 2020 Raomx. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.
// Package jwt is a easy and minimal implementation of JWT, and just implements HMAC SHA-256.
package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Set the registered claims.
func (claims Claims) setRegClaims() {
	claims[issuer] = iss
	claims[tokenID] = UUID()
	now := time.Now()

	claims[issuedAt] = now.Unix()
	claims[notBeforeAt] = now.Unix()
	claims[expiresAt] = now.Add(tokenDur).Unix()
}

//Validate registered claims
func (claims Claims) validate() error {

	now := time.Now().Unix()

	if err := claims.checkExpires(now, notBeforeAt); err != nil {
		return err
	}

	if err := claims.checkExpires(now, expiresAt); err != nil {
		return err
	}

	if err := claims.checkExpires(now, issuedAt); err != nil {
		return err
	}

	if !claims.Has(issuer) || claims[issuer] != iss {
		return errClaimValueInvalid
	}

	if !claims.Has(tokenID) {
		return errClaimValueInvalid
	}

	return nil
}

//Validate expires, include notBeforeAt, expiresAt and issuedAt.
func (claims Claims) checkExpires(now int64, tag string) error {
	if claims.Has(tag) {
		exp := claims[tag]
		var target int64
		switch val := exp.(type) {
		case float32:
			target = int64(val)
		case float64:
			target = int64(val)
		case int8:
			target = int64(val)
		case int16:
			target = int64(val)
		case int:
			target = int64(val)
		case int32:
			target = int64(val)
		case int64:
			target = val
		case uint8:
			target = int64(val)
		case uint16:
			target = int64(val)
		case uint32:
			target = int64(val)
		case uint:
			target = int64(val)
		case uint64:
			target = int64(val)
		default:
			return errClaimValueInvalid
		}
		if tag == notBeforeAt && now < target {
			return errClaimValueInvalid
		}
		if tag == expiresAt && now > target {
			return errClaimValueInvalid
		}
		if tag == issuedAt && now < target {
			return errClaimValueInvalid
		}
	} else {
		return errClaimValueInvalid
	}
	return nil
}

////Validate header
func verifyHeader(headToken string) error {

	var headClaims = make(map[string]string)
	headByte, err := base64.RawURLEncoding.DecodeString(headToken)

	if err != nil {
		return err
	}
	if err := json.Unmarshal(headByte, &headClaims); err != nil {
		return err
	}
	if val, ok := headClaims["alg"]; !ok || val != "HS256" {
		return errHeaderValueInvalid
	}

	if val, ok := headClaims["typ"]; !ok || val != "JWT" {
		return errHeaderValueInvalid
	}

	return nil
}

//Validate token secret.
func verifyToken(token string) bool {
	t := strings.Split(token, ".")
	if len(t) != 3 {
		return false
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(fmt.Sprintf("%s.%s", t[0], t[1])))
	expectedMAC := mac.Sum(nil)
	signature, err := base64.RawURLEncoding.DecodeString(t[2])
	if err != nil {
		log.Printf("Verify token %s: %v", token, err)
		return false
	}
	return hmac.Equal(signature, expectedMAC)
}

func setSecret(s string) {
	secret = []byte(s)
}
