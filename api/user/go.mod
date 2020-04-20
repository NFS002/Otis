module gitlab.com/otis-team/backend/api/user

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro v1.18.0
	gitlab.com/otis-team/backend/auth v0.0.0-20200410103937-7aa75954b834 // indirect
	gitlab.com/otis-team/backend/service/transaction v0.0.0-20200331103832-dcee07a226fa
	gitlab.com/otis-team/backend/service/user v0.0.0-20200327173517-c23610242505
)

// replace gitlab.com/otis-team/backend/auth => ../../auth