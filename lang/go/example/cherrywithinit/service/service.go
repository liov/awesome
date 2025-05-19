package service

import (
	"context"
	"strconv"
	user "test/example/cherrywithinit/proto"

	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/errors/errcode"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/hopeio/context/httpctx"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (u *UserService) Signup(ctx context.Context, req *user.SignupReq) (*wrapperspb.StringValue, error) {
	ctxi, _ := httpctx.FromContext(ctx)
	defer ctxi.StartSpanEnd("")()
	if req.Mail == "" && req.Phone == "" {
		return nil, errcode.InvalidArgument.Msg("请填写邮箱或手机号")
	}

	return &wrapperspb.StringValue{Value: "注册成功"}, nil
}

func Test(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	ctxi, _ := httpctx.FromContext(ctx.Request.Context())
	defer ctxi.StartSpanEnd("")()
	ctx.JSON(200, id)
}
