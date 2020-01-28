package simplestream_test

import (
	"testing"

	"github.com/arekmano/simplestream"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const testTableName = "TestTable-dev"

type dynamoStruct struct {
	User string `json:"user"`
	ID   string `json:"id"`
}

type invalidDynamoStruct struct {
	Value string
}

var region = "us-west-2"
var conf = aws.NewConfig().
	WithCredentials(credentials.NewEnvCredentials()).
	WithRegion("us-west-2")
var s = session.Must(session.NewSession(conf))
var db = dynamodb.New(s)
var dest = simplestream.DynamoDBDestination{
	DB:        db,
	TableName: testTableName,
}

func TestPut_valid(t *testing.T) {
	// Test data
	input := dynamoStruct{
		ID:   "test",
		User: "test",
	}

	// Execute
	err := dest.Put(input)

	// Verify
	if err != nil {
		t.Fatal(err)
	}
}

func TestPut_invalidInput(t *testing.T) {
	// Test data
	input := "Abc"

	// Execute
	err := dest.Put(input)

	// Verify
	if err == nil {
		t.Fatal("Did not throw error")
	}
}

func TestPut_invalidRequest(t *testing.T) {
	// Test data
	input := invalidDynamoStruct{
		Value: "test",
	}

	// Execute
	err := dest.Put(input)

	// Verify
	if err == nil {
		t.Fatal("Did not throw error")
	}
}
