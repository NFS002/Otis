module gitlab.com/otis-team/backend/service/transaction

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro v1.18.0
	gitlab.com/otis-team/backend/db v0.0.0-00010101000000-000000000000
	gitlab.com/otis-team/backend/service/merchant v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	gitlab.com/otis-team/backend/api/merchant => /Users/noah/Otis/backend/api/merchant/
	gitlab.com/otis-team/backend/api/user => /Users/noah/Otis/backend/api/user
	gitlab.com/otis-team/backend/db => /Users/noah/Otis/backend/db
	gitlab.com/otis-team/backend/service/merchant => /Users/noah/Otis/backend/service/merchant
	gitlab.com/otis-team/backend/service/transaction => /Users/noah/Otis/backend/service/transaction
	gitlab.com/otis-team/backend/service/user => /Users/noah/Otis/backend/service/user
)
