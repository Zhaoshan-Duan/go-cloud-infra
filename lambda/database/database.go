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

// InsertUserIfNotExists checks if a user exists, and if not, inserts the user into the database.
func (u DynamoDBClient) InsertUserIfNotExists(user types.RegisterUser) (bool, error) {
	// Check if user exists
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(user.Username)},
		},
	})

	if err != nil {
		return false, err
	}

	if result.Item != nil {
		// User already exists
		return false, nil
	}

	// User does not exist, insert new user
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(user.Username)},
			"password": {S: aws.String(user.Password)},
		},
	}
	_, err = u.databaseStore.PutItem(item)
	if err != nil {
		return false, err
	}
	return true, nil
}
