package models

type User struct {
	ID   string `dynamodbav:"id" json:"id"`
	Nome string `dynamodbav:"nome" json:"nome"`
}
