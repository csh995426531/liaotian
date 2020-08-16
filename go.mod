module liaotian

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/config/source/configmap/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	gopkg.in/yaml.v2 v2.2.8 // indirect
	k8s.io/api v0.17.1 // indirect
	k8s.io/apimachinery v0.17.1 // indirect
	k8s.io/client-go v11.0.0+incompatible // indirect
)

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.0
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
