module gitlab.com/otis_team/backend/db

go 1.14

require (
	github.com/aws/aws-sdk-go v1.30.4
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	gitlab.com/otis_team/backend/dtypes v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d // indirect
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace (
	gitlab.com/otis_team/backend/api/merchant => /Users/noah/Otis/backend/api/merchant/ //_LOCAL
	gitlab.com/otis_team/backend/api/user => /Users/noah/Otis/backend/api/user //_LOCAL
	gitlab.com/otis_team/backend/db => /Users/noah/Otis/backend/db //_LOCAL
	gitlab.com/otis_team/backend/dtypes => /Users/noah/Otis/backend/dtypes //_LOCAL
	gitlab.com/otis_team/backend/service/merchant => /Users/noah/Otis/backend/service/merchant //_LOCAL
	gitlab.com/otis_team/backend/service/transaction => /Users/noah/Otis/backend/service/transaction //_LOCAL
	gitlab.com/otis_team/backend/service/user => /Users/noah/Otis/backend/service/user //_LOCAL
)
