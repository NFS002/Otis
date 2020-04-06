module gitlab.com/otis-team/backend/service/transaction

go 1.13

require (
	github.com/aws/aws-sdk-go v1.30.4
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.3.0
	github.com/satori/go.uuid v1.2.0
	gitlab.com/otis-team/backend/db v0.0.0-00010101000000-000000000000 // indirect
	go.mongodb.org/mongo-driver v1.3.1
)

replace (
	gitlab.com/otis-team/backend/api/merchant => /Usets/noah/backend/api/user
	gitlab.com/otis-team/backend/api/user => /Users/noah/backend/api/user
	gitlab.com/otis-team/backend/db => /Users/noah/backend/db
	gitlab.com/otis-team/backend/service/merchant => /Users/noah/backend/service/merchant
	gitlab.com/otis-team/backend/service/transaction => /Users/noah/backend/service/transaction
	gitlab.com/otis-team/backend/service/user => /Users/noah/backend/service/user
)
