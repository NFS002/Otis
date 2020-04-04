package main

import (
	"flag"
	"log"
	"errors"
	"strings"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"gitlab.com/otis-team/backend/db/client"
	"gitlab.com/otis-team/backend/db/schema"
)

type dynamoClient client.DynamoClient

// Default flags
var (

	// "All" will include all tables for which a schema exists
	defaultNames = "User,Transaction,Merchant"

	defaultOperation = "create"

	namesToSchemaMap = map[string]*dynamodb.CreateTableInput {
		"transaction": schema.TransactionSchema,
		"user": schema.UserSchema,
		"merchant": schema.MerchantSchema, 
	}
)

func (c *dynamoClient) createTable(schema *dynamodb.CreateTableInput) error {
	var err error
	_, err = c.Client.CreateTable(schema)
	return err
}

func (c *dynamoClient) deleteTable(table string) error {
	var err error
	deleteTableInput := &dynamodb.DeleteTableInput {
		TableName: aws.String(table),
	}
	_, err = c.Client.DeleteTable(deleteTableInput)
	return err
}

func (c *dynamoClient) execute(table string, operation string) error {
	table = strings.ToLower(table)
	schema, exists := namesToSchemaMap[table]
	if !exists {
		return errors.New("Table '" + string(table) + "' could not be resolved to a valid schema")
	}
	var err error
	switch operation {
		case "create":
			err = c.createTable(schema)
		case "delete":
			err = c.deleteTable(table)
		case "recreate":
			err = c.deleteTable(table)
			if err != nil {
				return err
			}
			err = c.createTable(schema)
		default:
			return errors.New("Operation '" + string(operation) + "' could not be resolved to a valid operation")
	}
	return err
}


// Program to create, delete, recreate DynamoDB tables 
// using schemas defined in "gitlab.com/otis-team/backend/db/schema"
func main() {

	// A comma sepearted list of the names of tables to perform the selected operation on
	// e.g --names="merchant,transaction"
	names := *flag.String("names", defaultNames, "A comma sepearted list of the names of tables")
	
	// Possible values:
	//
	//  Create new tables with the given schemas, only if they do not already exist.
	//	operation="create"
	//
	//  Delete any existing tables and create new ones with the given schemas.
	//  Warning! This option will wipe any existing data!
	//	operation="recreate" 
	//
	//  Delete the given tables, partitions, and any data from the database.
	//	operation="delete"
	operation := *flag.String("operation",defaultOperation,"The operation to perform")
	flag.Parse()

	dynamoClient := &dynamoClient{}

	tables := strings.Split(names, ",")
	for _, table := range tables {
			err := dynamoClient.execute(table, operation)
			if err != nil {
				panic(err)
			}
	}
	log.Print("All operations executed successfully")
}