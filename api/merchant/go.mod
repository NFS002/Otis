module gitlab.com/otis-team/backend/api/merchant

go 1.13

replace github.com/coreos/etcd v3.3.17+incompatible => github.com/coreos/etcd v3.3.4+incompatible

require (
	github.com/go-log/log v0.2.0 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.3.0
	github.com/miekg/dns v1.1.29 // indirect
	gitlab.com/otis-team/backend/service/merchant v0.0.0-20200329213942-4509856c7460 // indirect
	gitlab.com/otis-team/backend/service/transaction v0.0.0-20200329213942-4509856c7460 // indirect
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
)
