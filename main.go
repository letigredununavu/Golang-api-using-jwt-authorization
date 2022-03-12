package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var port int = 8000

const registerPage string = `
<h1>Login</h1>
<form method="post" action="/addUser">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Register</button>
</form>
`

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

const indexPage string = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
<form method="post" action="/register">
	<button type="submit">Register</button>
</form>
`

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

	claims := customClaims{
		Username: player.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "MonAppDeCultutreG",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// La clef devrait être secrète pour vrai
	signedToken, err := token.SignedString([]byte("MaClefSecrete"))
	if err != nil {
		log.Println(err.Error())
		response.WriteHeader(http.StatusUnauthorized)
		return
	}
	http.SetCookie(
		response,
		&http.Cookie{
			Name:     "jwt_access_token",
			Value:    signedToken,
			Expires:  expirationTime,
			HttpOnly: true,
			Path:     "/",
		})

	http.Redirect(response, request, redirectTarget, 302)

}

const createQuestionPage string = `
<h1>Create Question</h1>
<form method="post" action="/addQuestion">
	<label for="body">Question</label>
	<input type="text" id="body" name="body">
	<label for="reponse">Reponse</label>
	<input type="text" id="reponse" name="reponse">
	<button type="submit">Create</button>
</form>
`

func createQuestionHandler(response http.ResponseWriter, request *http.Request) {
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
	fmt.Fprint(response, createQuestionPage)
}

func getDbPageHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	db, err := json.Marshal(getAllDB())

	if err != nil {
		log.Println(err)
	}

	response.Write(db)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/db", getDbPageHandler)
	router.HandleFunc("/create", createQuestionPageHandler)
	router.HandleFunc("/addQuestion", createQuestionHandler).Methods("POST")
	router.HandleFunc("/register", registerPageHandler)
	router.HandleFunc("/addUser", registerHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	fmt.Println("Bonsoir")

	http.Handle("/", router)
	log.Println("main: running simple server on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Could not start server", err)
	}
}
