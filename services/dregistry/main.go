package main

import (
	"log"

	micro "github.com/micro/go-micro"
	proto "github.com/srizzling/gotham/services/dregistry/proto"
	dregistry "github.com/srizzling/gotham/services/dregistry/src"
)

func main() {
	service := micro.NewService(
		micro.Name("DRegistry"),
	)

	proto.RegisterDRegistryHandler(service.Server(), &dregistry.DRegistry{
		Devices: dregistry.loadData("data/data.json"),
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
