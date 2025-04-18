package api

import (
	"github.com/hopeio/collection/cherrywithinit/proto"
	userService "github.com/hopeio/collection/cherrywithinit/service"
	"google.golang.org/grpc"
)

func GrpcRegister(gs *grpc.Server) {
	user.RegisterUserServiceServer(gs, userService.GetUserService())

}
