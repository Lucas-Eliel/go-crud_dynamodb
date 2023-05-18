package repositories

import (
	"context"
	"fmt"

	"crud_dynamodb/internal/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Repository interface {
	FindAll() []models.User
	FindById(id string) models.User
	Create(user models.User)
	Delete(id string)
}

type RepositoryImpl struct {
	db *dynamodb.Client
}

func NewRepository(db *dynamodb.Client) Repository {
	return RepositoryImpl{db}
}

// FindAll retorna todos os usuários cadastrados no banco de dados do DynamoDB
func (repo RepositoryImpl) FindAll() []models.User {

	out, err := repo.db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("user"),
	})

	if err != nil {
		panic(err)
	}

	var users []models.User

	err = attributevalue.UnmarshalListOfMaps(out.Items, &users)

	if err != nil {
		panic(err)
	}

	return users
}

// FindById retorna o usuário especifico pesquisado por ID
func (repo RepositoryImpl) FindById(id string) models.User {

	keyEx := expression.Key("id").Equal(expression.Value(id))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()

	out, err := repo.db.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String("user"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		panic(err)
	}

	var users []models.User

	err = attributevalue.UnmarshalListOfMaps(out.Items, &users)

	if err != nil {
		panic(err)
	}

	return users[0]
}

// Create faz o processo de criação de um usuário no banco de dados do dynamo db
func (repo RepositoryImpl) Create(user models.User) {

	fmt.Println(user.ID, user.Nome)

	userConverted, error := attributevalue.MarshalMap(user)

	fmt.Println(userConverted["id"])

	if error != nil {
		panic(error)
	}

	_, err := repo.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item:      userConverted,
	})

	if err != nil {
		panic(err)
	}
}

// Delete faz o processo de exclusão de um usuário no banco de dados do dynamo db
func (repo RepositoryImpl) Delete(id string) {

	userdb := repo.FindById(id)

	user, error := attributevalue.MarshalMap(userdb)

	if error != nil {
		panic(error)
	}

	_, err := repo.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("user"),
		Key:       user,
	})

	if err != nil {
		panic(err)
	}
}
