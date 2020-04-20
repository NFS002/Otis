package migrations

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// TransactionSchema : DynamoDB loose (NoSQL) migrations for the transaction table
var TransactionSchema *dynamodb.CreateTableInput = &dynamodb.CreateTableInput{

	// An array of attributes that describe the key migrations for the table and indexes.
    // AttributeDefinitions is a required field
	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		{

			// Required
			AttributeName: aws.String("transactionId"),

			// The data type for the attribute, where:
			//
			//    * S - the attribute is of type String
			//
			//    * N - the attribute is of type Number
			//
			//    * B - the attribute is of type Binary
			//
			// AttributeType is a required field
			AttributeType: aws.String("S"),
		},
		{
			// Required
			AttributeName: aws.String("merchantId"),

			// The data type for the attribute, where:
			//
			//    * S - the attribute is of type String
			//
			//    * N - the attribute is of type Number
			//
			//    * B - the attribute is of type Binary
			//
			// AttributeType is a required field
			AttributeType: aws.String("S"),
		},
	},

	KeySchema: []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("transactionId"),
			KeyType:       aws.String("HASH"),
		},
		{
			AttributeName: aws.String("merchantId"),
			KeyType:       aws.String("RANGE"),
		},
	},

	// For current minimum and maximum provisioned throughput values, see Limits
    // (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	// If using a PAY_PER_REQUEST billing model, this value is automatically set to
	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(0),
		WriteCapacityUnits: aws.Int64(0),
	},

	// The settings for DynamoDB Streams on the table. These settings consist of:
    //
    //    * StreamEnabled - Indicates whether DynamoDB Streams is to be enabled
    //    (true) or disabled (false).
    //
    //    * StreamViewType - When an item in the table is modified, StreamViewType
    //    determines what information is written to the table's stream. Valid values
    //    for StreamViewType are: KEYS_ONLY - Only the key attributes of the modified
    //    item are written to the stream. NEW_IMAGE - The entire item, as it appears
    //    after it was modified, is written to the stream. OLD_IMAGE - The entire
    //    item, as it appeared before it was modified, is written to the stream.
    //    NEW_AND_OLD_IMAGES - Both the new and the old item images of the item
    //    are written to the stream.
    StreamSpecification: &dynamodb.StreamSpecification {
		// Indicates whether DynamoDB Streams is enabled (true) or disabled (false)
		// on the table.
		//
		// StreamEnabled is a required field
		StreamEnabled: aws.Bool(false),

		// When an item in the table is modified, StreamViewType determines what information
		// is written to the stream for this table. Valid values for StreamViewType
		// are:
		//
		//    * KEYS_ONLY - Only the key attributes of the modified item are written
		//    to the stream.
		//
		//    * NEW_IMAGE - The entire item, as it appears after it was modified, is
		//    written to the stream.
		//
		//    * OLD_IMAGE - The entire item, as it appeared before it was modified,
		//    is written to the stream.
		//
		//    * NEW_AND_OLD_IMAGES - Both the new and the old item images of the item
		//    are written to the stream.
		StreamViewType: aws.String("NEW_IMAGE"),
	},

	// Table Name (required)
	TableName: aws.String("Transaction"),

	// Controls how you are charged for read and write throughput and how you manage
    // capacity. This setting can be changed later.
    //
	//    * PROVISIONED - For predictable workloads 
	//		https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
    //
	//    * PAY_PER_REQUEST - For unpredictabl workloads. 
	//		(https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
	//
	BillingMode: aws.String("PAY_PER_REQUEST"),


	// The settings used to enable server-side encryption.
	SSESpecification: &dynamodb.SSESpecification{
		
		// Indicates whether server-side encryption is done using an AWS managed CMK
		// or an AWS owned CMK. If enabled (true), server-side encryption type is set
		// to KMS and an AWS managed CMK is used (AWS KMS charges apply). If disabled
		// (false) or not specified, server-side encryption is set to AWS owned CMK.
		Enabled: aws.Bool(false),

		// The AWS KMS customer master key (CMK) that should be used for the AWS KMS
		// encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN),
		// alias name, or alias ARN. Note that you should only provide this parameter
		// if the key is different from the default DynamoDB customer master key alias/aws/dynamodb.
		//KMSMasterKeyId: aws.String("?"),

		// Server-side encryption type. The only supported value is:
		//
		//    * KMS - Server-side encryption that uses AWS Key Management Service. The
		//    key is stored in your account and is managed by AWS KMS (AWS KMS charges
		//    apply).
		SSEType: aws.String("KMS"),
	},

	// One or more local secondary indexes (the maximum is 5) to be created on the
    // table. Each index is scoped to a given partition key value. There is a 10
    // GB size limit per partition key value; otherwise, the size of a local secondary
    // index is unconstrained.
    //
    // Each local secondary index in the array includes the following:
    //
    //    * IndexName - The name of the local secondary index. Must be unique only
    //    for this table.
    //
    //    * KeySchema - Specifies the key migrations for the local secondary index.
    //    The key migrations must begin with the same partition key as the table.
    //
    //    * Projection - Specifies attributes that are copied (projected) from the
    //    table into the index. These are in addition to the primary key attributes
    //    and index key attributes, which are automatically projected. Each attribute
    //    specification is composed of: ProjectionType - One of the following: KEYS_ONLY
    //    - Only the index and primary keys are projected into the index. INCLUDE
    //    - Only the specified table attributes are projected into the index. The
    //    list of projected attributes is in NonKeyAttributes. ALL - All of the
    //    table attributes are projected into the index. NonKeyAttributes - A list
    //    of one or more non-key attribute names that are projected into the secondary
    //    index. The total count of attributes provided in NonKeyAttributes, summed
    //    across all of the secondary indexes, must not exceed 100. If you project
    //    the same attribute into two different indexes, this counts as two distinct
    //    attributes when determining the total.
    LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex {
		&dynamodb.LocalSecondaryIndex{
			IndexName: aws.String("date"),

			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("date"),
					KeyType:       aws.String("S"),
				},
			},

			Projection: &dynamodb.Projection {

				// Represents the non-key attribute names which will be projected into the index.
				//
				// For local secondary indexes, the total count of NonKeyAttributes summed across
				// all of the local secondary indexes, must not exceed 20. If you project the
				// same attribute into two different indexes, this counts as two distinct attributes
				// when determining the total.
    			/*NonKeyAttributes: []*string{
					aws.String("?"),
					aws.String("?"),
				},*/

				// The set of attributes that are projected into the index:
				//
				//    * KEYS_ONLY - Only the index and primary keys are projected into the index.
				//
				//    * INCLUDE - Only the specified table attributes are projected into the
				//    index. The list of projected attributes is in NonKeyAttributes.
				//
				//    * ALL - All of the table attributes are projected into the index.
				ProjectionType: aws.String("ALL"),
			},
		},
	},


	// One or more global secondary indexes (the maximum is 20) to be created on
    // the table. Each global secondary index in the array includes the following:
    //
    //    * IndexName - The name of the global secondary index. Must be unique only
    //    for this table.
    //
    //    * KeySchema - Specifies the key migrations for the global secondary index.
    //
    //    * Projection - Specifies attributes that are copied (projected) from the
    //    table into the index. These are in addition to the primary key attributes
    //    and index key attributes, which are automatically projected. Each attribute
    //    specification is composed of: ProjectionType - One of the following: KEYS_ONLY
    //    - Only the index and primary keys are projected into the index. INCLUDE
    //    - Only the specified table attributes are projected into the index. The
    //    list of projected attributes is in NonKeyAttributes. ALL - All of the
    //    table attributes are projected into the index. NonKeyAttributes - A list
    //    of one or more non-key attribute names that are projected into the secondary
    //    index. The total count of attributes provided in NonKeyAttributes, summed
    //    across all of the secondary indexes, must not exceed 100. If you project
    //    the same attribute into two different indexes, this counts as two distinct
    //    attributes when determining the total.
    //
    //    * ProvisionedThroughput - The provisioned throughput settings for the
    //    global secondary index, consisting of read and write capacity units.
	/*GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex {
		&dynamodb.GlobalSecondaryIndex{
			IndexName: aws.String("?"),

			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("?"),
					KeyType:       aws.String("?"),
				},
				{
					AttributeName: aws.String("?"),
					KeyType:       aws.String("?"),
				},
			},

			// For current minimum and maximum provisioned throughput values, see Limits
			// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
			// in the Amazon DynamoDB Developer Guide.
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(10),
			},

			Projection: &dynamodb.Projection {

				// Represents the non-key attribute names which will be projected into the index.
				//
				// For local secondary indexes, the total count of NonKeyAttributes summed across
				// all of the local secondary indexes, must not exceed 20. If you project the
				// same attribute into two different indexes, this counts as two distinct attributes
				// when determining the total.
    			NonKeyAttributes: []*string{
					aws.String("?"),
					aws.String("?"),
				},

				// The set of attributes that are projected into the index:
				//
				//    * KEYS_ONLY - Only the index and primary keys are projected into the index.
				//
				//    * INCLUDE - Only the specified table attributes are projected into the
				//    index. The list of projected attributes is in NonKeyAttributes.
				//
				//    * ALL - All of the table attributes are projected into the index.
				ProjectionType: aws.String(""),
			},
		},
	},*/

	// A list of key-value pairs to label the table. For more information, see Tagging
    // https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html).
    /*Tags: []*dynamodb.Tag {
		&dynamodb.Tag{
			Key: aws.String("key"),
			Value: aws.String("value"),
		},
	},*/
}