package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "UserTable"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	return DynamoDBClient{
		databaseStore: dynamodb.New(session.Must(session.NewSession())),
	}
}

// Does this user exist in the database?
func (u DynamoDBClient) DoesUserExist(username string) (bool, error) {
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(username)},
		},
	})

	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

// How do I insert a new record into the database?
func (u DynamoDBClient) InsertUser(user types.RegisterUser) error {
	// assemble the item
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(user.Username)},
			"password": {S: aws.String(user.Password)},
		},
	}
	// insert
	_, err := u.databaseStore.PutItem(item)
	if err != nil {
		return err
	}
	return nil
}
