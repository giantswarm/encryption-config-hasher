module github.com/giantswarm/encryption-config-hasher

go 1.16

require (
	golang.org/x/crypto v0.14.0
	k8s.io/api v0.25.2
	k8s.io/apimachinery v0.25.2
	k8s.io/client-go v0.25.2
	sigs.k8s.io/controller-runtime v0.13.0
)

replace golang.org/x/net => golang.org/x/net v0.17.0
