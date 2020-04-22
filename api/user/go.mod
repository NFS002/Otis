module gitlab.com/otis_team/backend/api/user

go 1.13

replace (
	gitlab.com/otis_team/backend/api/merchant => /Users/noah/Otis/backend/api/merchant/ //_LOCAL
	gitlab.com/otis_team/backend/api/user => /Users/noah/Otis/backend/api/user //_LOCAL
	gitlab.com/otis_team/backend/db => /Users/noah/Otis/backend/db //_LOCAL
	gitlab.com/otis_team/backend/dtypes => /Users/noah/Otis/backend/dtypes //_LOCAL
	gitlab.com/otis_team/backend/service/merchant => /Users/noah/Otis/backend/service/merchant //_LOCAL
	gitlab.com/otis_team/backend/service/transaction => /Users/noah/Otis/backend/service/transaction //_LOCAL
	gitlab.com/otis_team/backend/service/user => /Users/noah/Otis/backend/service/user //_LOCAL
)

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro/v2 v2.4.0
	gitlab.com/otis_team/backend/db v0.0.0-00010101000000-000000000000 // indirect
)
