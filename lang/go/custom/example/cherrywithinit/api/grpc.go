package api

import (
	"test/example/cherrywithinit/proto"
	userService "test/example/cherrywithinit/service"
	"google.golang.org/grpc"
)

func GrpcRegister(gs *grpc.Server) {
	user.RegisterUserServiceServer(gs, userService.GetUserService())

}
