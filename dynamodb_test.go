package simplestream_test

import (
	"testing"

	"github.com/arekmano/simplestream"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const testTableName = "go-adventures-TestTable-dev"

type dynamoStruct struct {
	ID    string
	Value string
}

type invalidDynamoStruct struct {
	Value string
}

var region = "us-west-2"
var s = session.Must(session.NewSession(&aws.Config{Region: &region}))
var db = dynamodb.New(s)
var dest = simplestream.DynamoDBDestination{
	DB:        db,
	TableName: testTableName,
}

func TestPut_valid(t *testing.T) {
	// Test data
	input := dynamoStruct{
		ID:    "test",
		Value: "test",
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
