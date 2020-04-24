module gitlab.com/otis_team/backend/service/user

go 1.13

replace (
	gitlab.com/otis_team/backend/api/merchant => ../../api/merchant/ //_LOCAL
	gitlab.com/otis_team/backend/api/user => ../../api/user //_LOCAL
	gitlab.com/otis_team/backend/db => ../../db //_LOCAL
	gitlab.com/otis_team/backend/dtypes => ../../dtypes //_LOCAL
	gitlab.com/otis_team/backend/service/merchant => ../../service/merchant //_LOCAL
	gitlab.com/otis_team/backend/service/transaction => ../../service/transaction //_LOCAL
	gitlab.com/otis_team/backend/service/user => ../../service/user //_LOCAL
)

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro v1.18.0
	gitlab.com/otis_team/backend/db v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/otis_team/backend/dtypes v0.0.0-00010101000000-000000000000
)
