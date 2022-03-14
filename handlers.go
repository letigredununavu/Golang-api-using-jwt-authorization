package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getDbPageHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	db, err := json.Marshal(getAllDB())

	if err != nil {
		log.Println(err)
	}

	response.Write(db)
}

func createQuestionHandler(response http.ResponseWriter, request *http.Request) {
	// Devrait probablement aller dans un middleware ou un fonction appart
	// -----------------
	cookie, err := request.Cookie("jwt_access_token")
	if err != nil {
		log.Println(err)
		response.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value
	ok := tokenValid(tokenString)
	if ok != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}
	//-------------------

	body := request.FormValue("body")
	reponse := request.FormValue("reponse")
	redirectTarget := "/db"
	if body != "" && reponse != "" {
		q := makeQuestion(body, reponse)
		addQuestion(q)
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func createQuestionPageHandler(response http.ResponseWriter, request *http.Request) {
	ip, err := getIP(request)
	if err != nil {
		log.Println(err)
	}
	log.Println(ip)
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(response, ip)
	fmt.Fprint(response, createQuestionPage)
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")

	redirectTarget := "/"

	// Vérifie que le joueur existe
	player, err := verifyPlayer(username, password)
	if err != nil {
		log.Println(err.Error())
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Si le joueur existe crée le cookie jwt avec 7min d'expiration
	expirationTime := time.Now().Add(time.Minute * 7)
	signedToken, err := createJWTToken(player, expirationTime)
	if err != nil {
		log.Println(err)
	}
	http.SetCookie(
		response,
		&http.Cookie{
			Name:     "jwt_access_token",
			Value:    *signedToken,
			Expires:  expirationTime,
			HttpOnly: true,
			Path:     "/",
		},
	)

	http.Redirect(response, request, redirectTarget, 302)

}

func registerPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, registerPage)
}

func registerHandler(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")
	redirectTarget := "/"
	// Faire les vérifications nécessaires en vraison
	player := createPlayer(username, password)
	addPlayer(player)
	log.Println(players)
	http.Redirect(response, request, redirectTarget, 302)
}

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	// Check si y'a un cookie, si oui l'imprime, sinon imprime l'erreur
	cookie, err := request.Cookie("jwt_access_token")
	if err != nil {
		log.Println("Dans le indexPageHandler")
		log.Println(err)
	}
	fmt.Println(cookie)
	fmt.Fprintf(response, indexPage)
}
