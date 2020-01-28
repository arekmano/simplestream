package simplestream

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sirupsen/logrus"
)

type DynamoDBDestination struct {
	DB        *dynamodb.DynamoDB
	TableName string
}

func (d *DynamoDBDestination) Put(p interface{}) error {
	logrus.
		WithField("item", p).
		Debug("Loading item into DynamoDB")
	av, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: &d.TableName,
		Item:      av,
	}
	logrus.
		WithField("request", input).
		Debug("PutItem")
	out, err := d.DB.PutItem(input)
	if err != nil {
		return err
	}
	logrus.
		WithField("output", out).
		Debug("Successfully loaded item into DynamoDB")
	return nil
}
