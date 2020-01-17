package jwt

import "errors"

//There are only four errors.
var (
	errClaimValueInvalid = errors.New("Claim value invalid")

	errTokenInvalid = errors.New("Token is invalid")

	errTokenHasExpired = errors.New("Token has expired")

	errHeaderValueInvalid = errors.New("Head value invalid")
)
