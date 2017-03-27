package main

import (
	"log"

	micro "github.com/micro/go-micro"
	proto "github.com/srizzling/gotham/services/base/proto"
	wol "github.com/srizzling/gotham/services/wol/src"
)

func main() {
	service := micro.NewService(
		micro.Name("gotham.services.WolService"),
	)

	proto.RegisterServiceHandler(service.Server(), new(wol.Wol))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
