[![Build Status](https://travis-ci.org/raomx/jwt.svg?branch=master)](https://travis-ci.org/raomx/jwt) [![Go Report Card](https://goreportcard.com/badge/github.com/raomx/jwt)](https://goreportcard.com/report/github.com/raomx/jwt) [![codecov.io](https://codecov.io/github/raomx/jwt/branch/master/graph/badge.svg)](https://codecov.io/github/raomx/jwt) [![GoDoc](https://godoc.org/github.com/robbert229/jwt?status.svg)](https://godoc.org/github.com/raomx/jwt)


# JWT
The JWT is a easy and minimal implementation of JWT, and just implements HMAC SHA-256 by go(golang).

## How to use:

### Get a jwt token
    claims := Claims {
        "name": raomx,
        "age":  38,
    }
    token := claims.Sign()


### Get Claims from token
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzkxOTc0OTUsImlhdCI6MTU3OTE5MDI5NSwiaXNzIjoiYXV0aC5leGFtcGxlLmNvbSIsImp0aSI6IjAxZTZjNTczLTQ4YzQtNDYyMi04M2U3LThiNjRhZDNkZjg0NyIsIm5iZiI6MTU3OTE5MDI5NSwibmFtZSI6InJhb214IiwiYWdlIjozOH0.3jGXEPaXLuUsH8R-m-BDQght3-IhoUHDO7kK5gR0CsA"
    claims, err :=  Verify(token)
    if err != nil {
        return fmt.Errorf("Verify %s err: %w", token, err)
    }
    name := claims["name"]

## What did the JWT do?
The JWT has just two APIs: Sign() and Parse().  
In Sign, The JWT sets tokenID, issuer, issuedAt, expiresAt, notBeforeAt and secret.  
In Verify, The JWT verifies token secret, tokenID, issuer, issuedAt, expiresAt, notBeforeAt and header.  
The secret is a list byte, length between 25-32.  

## What can you do?
You can change expires duration by SettokenDur() which default is 2 hours.  
You can change issuer by SetIss().  
And you can add sub, aud and anything you want by Claims.  
