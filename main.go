package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port int = 8000

const indexPage string = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

const createQuestionPage string = `
<h1>Create Question</h1>
<form method="post" action="/addQuestion">
	<label for="question">Question</label>
	<input type="text" id="question" name="question">
	<label for="reponse">Reponse</label>
	<input type="text" id="reponse" name="reponse">
	<button type="submit">Create</button>
</form>
`

func createQuestionHandler(response http.ResponseWriter, request *http.Request) {
	body := request.FormValue("question")
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

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
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
	fmt.Println("Bonsoir")
	question := makeQuestion("Quel arbre fait le kakis ?", "Le plaqueminier")
	addQuestion(question)

	hola()

	http.Handle("/", router)
	log.Println("main: running simple server on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Could not start server", err)
	}
}