module github.com/sputnik-systems/backups-operator

go 1.16

require (
	github.com/aws/aws-sdk-go v1.41.2
	github.com/hasura/go-graphql-client v0.3.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/sputnik-systems/backups-storage v0.0.0-20211013190640-c9ca413c45ad
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	sigs.k8s.io/controller-runtime v0.9.2
)
