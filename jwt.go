package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string = "MaClefQuiDevraitEtrePlusLongueEtSecreteQueCa"

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return token, nil
}

func tokenValid(tokenString string) error {
	token, err := verifyToken(tokenString)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}
	return nil
}

func createJWTToken(player *Player, expirationTime time.Time) (*string, error) {
	claims := customClaims{
		Username: player.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    SECRET_KEY,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// La clef devrait être secrète pour vrai
	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}
