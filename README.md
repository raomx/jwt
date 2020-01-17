[![GoDoc](https://godoc.org/github.com/robbert229/jwt?status.svg)](https://godoc.org/github.com/raomx/jwt)


# JWT
The JWT is a easy and minimal implementation of JWT, and just implements HMAC SHA-256.

## How to use:

### Get a jwt token
    claims = Claims {
        "name": raomx,
        "age":  38,
    }
    token := claims.GetToken()


### Get Claims from token
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzkxOTc0OTUsImlhdCI6MTU3OTE5MDI5NSwiaXNzIjoiYXV0aC5leGFtcGxlLmNvbSIsImp0aSI6IjAxZTZjNTczLTQ4YzQtNDYyMi04M2U3LThiNjRhZDNkZjg0NyIsIm5iZiI6MTU3OTE5MDI5NSwibmFtZSI6InJhb214IiwiYWdlIjozOH0.3jGXEPaXLuUsH8R-m-BDQght3-IhoUHDO7kK5gR0CsA"
    claims, err :=  Parse(token)
    if err := nil {
        return fmt.Errorf("Parse %s err: %w", token, err)
    }
    name := claims["name"]
