package processor

import "github.com/gin-gonic/gin"

type IHandlerManager interface {
	Engine() *gin.Engine
}
