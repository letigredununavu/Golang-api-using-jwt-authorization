package main

import (
	"errors"
)

type Question struct {
	Id      uint32 `json:"id"`
	Vote    int32  `json:"vote"`
	Body    string `json:"body"`
	Reponse string `json:"reponse"`
}

var ErrorQuestionIdNotFound error = errors.New("Question Id not found")

var serialQuestionId uint32 = 1

var db []Question

func makeQuestion(body string, reponse string) *Question {
	question := &Question{
		Id:      serialQuestionId,
		Vote:    0,
		Body:    body,
		Reponse: reponse,
	}
	serialQuestionId += 1

	return question
}

func addQuestion(question *Question) int {
	db = append(db, *question)
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
	return ErrorQuestionIdNotFound
}

func updateQuestion(id uint32, question Question) error {
	for i, question := range db {
		if question.Id == id {
			db[i] = question
			return nil
		}
	}
	return ErrorQuestionIdNotFound
}

func getQuestionById(id uint32) (*Question, error) {
	for _, question := range db {
		if question.Id == id {
			return &question, nil
		}
	}

	return nil, ErrorQuestionIdNotFound

}

func getAllDB() []Question {
	return db
}
