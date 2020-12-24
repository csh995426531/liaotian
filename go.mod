module liaotian

go 1.14

require (
	github.com/SkyAPM/go2sky v0.6.0
	github.com/SkyAPM/go2sky-plugins/micro v0.0.0-20201013142527-b9cca04e7aa8
	github.com/gin-gonic/gin v1.6.3
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.5 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	github.com/micro/go-plugins v1.5.1
	github.com/miekg/dns v1.1.35 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/onsi/ginkgo v1.13.0 // indirect
	github.com/prometheus/common v0.7.0
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/ugorji/go v1.1.9 // indirect
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
	golang.org/x/sys v0.0.0-20201118182958-a01c418693c7 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20200812195022-5ae4c3c160a0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20201117123952-62d171c70ae1 // indirect
	google.golang.org/grpc v1.33.2 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.7
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/api v0.17.1 // indirect
	k8s.io/apimachinery v0.17.1 // indirect
	k8s.io/utils v0.0.0-20200109141947-94aeca20bf09 // indirect
)

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.0
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
