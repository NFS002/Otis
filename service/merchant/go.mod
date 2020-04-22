module gitlab.com/otis_team/backend/service/merchant

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

replace github.com/coreos/etcd v3.3.17+incompatible => github.com/coreos/etcd v3.3.4+incompatible

require (
)
