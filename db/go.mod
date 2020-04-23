module gitlab.com/otis_team/backend/db

go 1.14

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
	github.com/aws/aws-sdk-go v1.30.12
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/jinzhu/gorm v1.9.12
	gitlab.com/otis_team/backend/dtypes v0.0.0-00010101000000-000000000000
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
)
