module gitlab.com/otis_team/backend/auth

go 1.14

require (
	github.com/auth0/go-jwt-middleware v0.0.0-20190805220309-36081240882b
	github.com/codegangsta/negroni v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/micro/go-micro v1.18.0
	gitlab.com/otis_team/backend/api/merchant v0.0.0-20200408150416-63532e1b178b
)

replace (
//	gitlab.com/otis_team/backend/api/merchant => ../../api/merchant/ //_LOCAL
//	gitlab.com/otis_team/backend/api/user => ../../api/user //_LOCAL
//	gitlab.com/otis_team/backend/db => ../../db //_LOCAL
//	gitlab.com/otis_team/backend/dtypes => ../../dtypes //_LOCAL
//	gitlab.com/otis_team/backend/service/merchant => ../../service/merchant //_LOCAL
//	gitlab.com/otis_team/backend/service/transaction => ../../service/transaction //_LOCAL
//	gitlab.com/otis_team/backend/service/user => ../../service/user //_LOCAL
)
