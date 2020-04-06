module gitlab.com/otis-team/backend/service/user

go 1.13

require (
	github.com/aws/aws-sdk-go v1.30.4
	github.com/go-log/log v0.2.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro v1.18.0
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/miekg/dns v1.1.29 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0
	gitlab.com/otis-team/backend/db v0.0.0-00010101000000-000000000000
	gitlab.com/otis-team/backend/service/merchant v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/otis-team/backend/service/transaction v0.0.0-00010101000000-000000000000 // indirect
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
)

replace (
	gitlab.com/otis-team/backend/api/merchant => /Usets/noah/backend/api/user
	gitlab.com/otis-team/backend/api/user => /Users/noah/backend/api/user
	gitlab.com/otis-team/backend/db => /Users/noah/backend/db
	gitlab.com/otis-team/backend/service/merchant => /Users/noah/backend/service/merchant
	gitlab.com/otis-team/backend/service/transaction => /Users/noah/backend/service/transaction
	gitlab.com/otis-team/backend/service/user => /Users/noah/backend/service/user
)
