module gitlab.com/otis-team/backend/service/merchant

go 1.13

require (
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.3.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/btree v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/lucas-clemente/quic-go v0.14.1 // indirect
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/miekg/dns v1.1.27 // indirect
	github.com/nats-io/nats-server/v2 v2.1.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	go.mongodb.org/mongo-driver v1.3.1
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/tools v0.0.0-20191216173652-a0e659d51361 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1 // indirect
	google.golang.org/grpc v1.26.0 // indirect
)

replace (
    gitlab.com/otis-team/backend/db => /Users/noah/backend/db
    gitlab.com/otis-team/backend/api/user => /Users/noah/backend/api/user
    gitlab.com/otis-team/backend/api/merchant => /Usets/noah/backend/api/user
    gitlab.com/otis-team/backend/service/user => /Users/noah/backend/service/user
    gitlab.com/otis-team/backend/service/transaction => /Users/noah/backend/service/transaction
    gitlab.com/otis-team/backend/service/merchant => /Users/noah/backend/service/merchant
)
