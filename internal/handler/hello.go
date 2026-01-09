package handler

import (
	"gin-hello-world/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	helloService service.HelloService // 依赖接口
}

func NewHelloHandler(helloService service.HelloService) *HelloHandler {
	return &HelloHandler{
		helloService: helloService,
	}
}

func (h *HelloHandler) Hello(c *gin.Context) {
	message := h.helloService.GetHelloMessage()
	demo, err := h.helloService.TestDemo()
	if err != nil {
		gin.Logger()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    message + demo,
	})
}
