package main

import (
	"errors"
	"fmt"
)

type Question struct {
	Id      uint32 `json:"id"`
	Vote    int32  `json:"vote"`
	Body    string `json:"body"`
	Reponse string `json:"reponse"`
}

var ErrorIdNotFound error = errors.New("Question Id not found")

var serialId uint32 = 1

var db []Question

func hola() {
	fmt.Println("Allo dans test.go")
}

func makeQuestion(body string, reponse string) Question {
	question := Question{
		Id:      serialId,
		Vote:    0,
		Body:    body,
		Reponse: reponse,
	}
	serialId += 1

	return question
}

func addQuestion(question Question) int {
	db = append(db, question)
	return 0
}

func removeQuestion(id uint32) error {
	for i, question := range db {
		if question.Id == id {
			// Syntax normal pour enlever l'élément i d'un tableau
			db = append(db[:i], db[i+1:]...)
			return nil
		}
	}
	return ErrorIdNotFound
}

func updateQuestion(id uint32, question Question) error {
	for i, question := range db {
		if question.Id == id {
			db[i] = question
			return nil
		}
	}
	return ErrorIdNotFound
}

func getQuestionById(id uint32) (Question, error) {
	for i, question := range db {
		if question.Id == id {
			return db[i], nil
		}
	}

	return Question{0, 0, "", ""}, ErrorIdNotFound

}

func getAllDB() []Question {
	return db
}