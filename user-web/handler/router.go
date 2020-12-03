package handler

import (
	"liaotian/middlewares/wrapper/skywalking/gin2micro"

	"github.com/SkyAPM/go2sky"
	// sky2gin "github.com/SkyAPM/go2sky-plugins/gin/v2"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	// "liaotian/plugins/wrapper/tracer/opentracing/gin2micro"
)

func InitRouters() *gin.Engine {
	// gin2micro.SetSamplingFrequency(50)
	report, err := reporter.NewGRPCReporter("oap.skywalking.svc.cluster.local:11800")
	if err != nil {
		log.Fatalf("crate grpc reporter error: %v \n", err)
	}
	tracer, err := go2sky.NewTracer("user-web", go2sky.WithReporter(report))
	if err != nil {
		log.Fatalf("crate tracer error: %v \n", err)
	} else {
		log.Infof("create trace oap.skywalking:11800 - user-web")
	}

	router := gin.Default()
	router.Use(gin2micro.Middleware(router, tracer))
	router.POST("login", Login)

	return router
}
