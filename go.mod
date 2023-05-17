module github.com/giantswarm/encryption-config-hasher

go 1.16

require (
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	k8s.io/api v0.17.17
	k8s.io/apimachinery v0.17.17
	k8s.io/client-go v0.17.17
	sigs.k8s.io/controller-runtime v0.5.14
)

replace (
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.0 => github.com/gorilla/websocket v1.4.2
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.12.2
	go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/net => golang.org/x/net v0.10.0
	golang.org/x/text => golang.org/x/text v0.9.0
)
