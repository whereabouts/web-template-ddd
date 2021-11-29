package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	valication "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/whereabouts/sdk/httpserver/handler"
	"github.com/whereabouts/sdk/logger"
	"github.com/whereabouts/web-template-ddd/application"
	"github.com/whereabouts/web-template-ddd/server/proto"
)

type TestHandler struct {
	logger      *logger.Entry
	application *application.TestApplication
}

// Init implements handler.
func (h *TestHandler) Init(router *gin.RouterGroup) {
	h.logger = logger.WithField("handler", "test")
	h.application = application.NewTestApplication()

	// registry http handler
	if router != nil {
		test := router.Group("test")
		test.GET("/ping", handler.New(h.Ping))
		test.POST("/register", handler.NewWithOptions(h.TestRegister))
	}

}

func (h *TestHandler) Ping(ctx context.Context, _ *struct{}) (*string, error) {
	resp := "pong"
	return &resp, nil
}

func (h *TestHandler) TestRegister(ctx context.Context, req *proto.TestRegisterReq) (*proto.TestRegisterResp, error) {
	// handler中通常需要做的事情:
	// 1.校验请求参数
	// 2.调用application, 响应其返回的数据
	// 3.很少的时候也会多个application协作

	if err := valication.ValidateStruct(req,
		valication.Field(&req.Name, valication.Required),
		valication.Field(&req.Password, valication.Required),
	); err != nil {
		return nil, err
	}

	test, err := h.application.TestRegister(ctx, req.Name, req.Password)
	if err != nil {
		return nil, err
	}

	return &proto.TestRegisterResp{
		Test: test,
	}, nil
}