package main

import (
	"github.com/dawn1806/common"
	"github.com/dawn1806/order/domain/repository"
	service2 "github.com/dawn1806/order/domain/service"
	"github.com/dawn1806/order/handler"
	order "github.com/dawn1806/order/proto/order"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var (
	QPS = 1000
)

func main() {

	//t, io, err := common.NewTracer("go.micro.service.order", "localhost:6831")
	//if err != nil {
	//	log.Error(err)
	//}
	//defer io.Close()
	//opentracing.SetGlobalTracer(t)

	db, err := gorm.Open("mysql", common.MysqlConnection)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// 第一次运行时创建表
	//repository.NewOrderRepository(db).InitTable()

	orderService := service2.NewOrderDataService(repository.NewOrderRepository(db))

	//common.PrometheusBoot(9092)

	// New Service
	service := micro.NewService(
		micro.Name("micro.order"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8003"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("127.0.0.1:2379"))),
		//micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		//micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), &handler.Order{OrderDataService: orderService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
