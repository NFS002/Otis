module gitlab.com/otis-team/backend/api/merchant

go 1.14

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.18.0
    gitlab.com/otis-team/backend/authentication v0.0.0-00010101000000-000000000000
	gitlab.com/otis-team/backend/service/merchant v0.0.0-20200330224117-cbbfc773f1a3
	gitlab.com/otis-team/backend/service/transaction v0.0.0-20200330224117-cbbfc773f1a3
)

replace gitlab.com/otis-team/backend/authentication => /Users/cjwilliams/gitlab/otis/backend/authentication //_LOCAL
