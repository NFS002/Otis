package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"encoding/json"
)

const (
	ApiIdentifier = "https://otis-app.com"
	Domain        = "https://otis-app.eu.auth0.com/"
)

var (
	SigningMethod = jwt.SigningMethodRS256
)

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N string `json:"n"`
	E string `json:"e"`
	X5c []string `json:"x5c"`
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(Domain + ".well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}

func validationKeyGetter(token *jwt.Token) (interface{}, error) {
	// Verify 'aud' claim
	aud := ApiIdentifier
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return token, errors.New("Invalid audience.")
	}
	// Verify 'iss' claim
	iss := Domain
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return token, errors.New("Invalid issuer.")
	}

	cert, err := getPemCert(token)
	if err != nil {
		panic(err.Error())
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func CheckAuthHeader(header string) error{
	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Fields(header)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return errors.New("Authorization header format must be Bearer {token}")
	}

	token := authHeaderParts[1]

	if token == "" {
		return errors.New("Required authorization token not found")
	}

	// Now parse the token
	parsedToken, err := jwt.Parse(token, validationKeyGetter)
	if err != nil {
		return fmt.Errorf("Error parsing token: %v", err)
	}

	if SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			SigningMethod.Alg(),
			parsedToken.Header["alg"])
		return fmt.Errorf("Error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		return errors.New("Token is invalid")
	}

	return nil
}