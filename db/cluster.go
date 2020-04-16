package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"gitlab.com/otis-team/backend/db/client"
	"os"
)

type rdsClient struct {
	*client.RDSClient
}


func (c* rdsClient) createCluster() (*rds.CreateDBClusterOutput, error) {
	input := &rds.CreateDBClusterInput{
		AvailabilityZones: []*string{
			aws.String(os.Getenv("DB_CLUSTER_ZONE")),
		},
		BackupRetentionPeriod:       aws.Int64(1),
		DBClusterIdentifier:         aws.String(os.Getenv("DB_CLUSTER_NAME")),
		DBClusterParameterGroupName: aws.String("mydbclusterparametergroup"),
		DatabaseName:                aws.String(os.Getenv("DB_CLUSTER_ZONE")),
		Engine:                      aws.String("aurora"),
		EngineVersion:               aws.String("5.6.10a"),
		MasterUserPassword:          aws.String("mypassword"),
		MasterUsername:              aws.String("myuser"),
		Port:                        aws.Int64(3306),
		StorageEncrypted:            aws.Bool(true),
	}

	return c.Client.CreateDBCluster(input)
}

func main() {
	cli := client.RDSClient{}
	err := cli.Init()
	if err != nil {
		panic(err)
	}
	rdCli := rdsClient{ &cli }
	result, err := rdCli.createCluster()
	fmt.Println(result, err)
	if err != nil {
		panic(err)
	}
}
