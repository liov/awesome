package service

import (
	"context"
	"github.com/gin-gonic/gin"
	user "github.com/hopeio/example/cherrywithinit/proto"
	"github.com/hopeio/utils/errors/errcode"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strconv"

	"github.com/hopeio/context/httpctx"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (u *UserService) Signup(ctx context.Context, req *user.SignupReq) (*wrapperspb.StringValue, error) {
	ctxi, _ := httpctx.FromContextValue(ctx)
	defer ctxi.StartSpanEnd("")()
	if req.Mail == "" && req.Phone == "" {
		return nil, errcode.InvalidArgument.Msg("请填写邮箱或手机号")
	}

	return &wrapperspb.StringValue{Value: "注册成功"}, nil
}

func Test(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	ctxi, _ := httpctx.FromContextValue(ctx.Request.Context())
	defer ctxi.StartSpanEnd("")()
	ctx.JSON(200, id)
}
